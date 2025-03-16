package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// 假设的用户数据（实际应用中应从数据库获取）
var users = map[string]string{
	"user1": "password123",
	"user2": "secret456",
}

// JWT密钥（实际应用中应从安全的地方获取）
var jwtSecret = []byte("your-secret-key")

// 用于在context中存储用户 ID 的 key
type contextKey string

const userIDKey contextKey = "userID"

// ErrorResponse 结构体，用于返回 JSON 错误
type ErrorResponse struct {
	Message string `json:"message"`
}

// APIError 结构体，更详细的错误信息，可以扩展
type APIError struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// writeJSONError 辅助函数，发送 JSON 格式的错误响应
func writeJSONError(w http.ResponseWriter, err APIError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Status)
	json.NewEncoder(w).Encode(err) // 直接将 err 结构体编码为 JSON，并写入响应
}

// loggingMiddleware 日志中间件
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

// authMiddleware JWT 鉴权中间件
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 获取 Authorization 头
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			writeJSONError(w, APIError{Status: http.StatusUnauthorized, Code: "missing_token", Message: "Missing authorization token"})
			return
		}
		// 验证 Bearer Token 格式
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			writeJSONError(w, APIError{Status: http.StatusUnauthorized, Code: "invalid_token", Message: "Invalid authorization token"})
			return
		}
		tokenString := parts[1]
		// 解析 JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 验证签名方法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})
		if err != nil {
			// 如果是 Token 过期错误，返回特定的错误码
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorExpired != 0 {
					writeJSONError(w, APIError{Status: http.StatusUnauthorized, Code: "token_expired", Message: "Token has expired"})
					return
				}
			}
			writeJSONError(w, APIError{Status: http.StatusUnauthorized, Code: "invalid_token", Message: "Invalid authorization token"})
			return
		}
		// 验证通过，提取 claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// 将用户 ID 存入 context
			userID, ok := claims["sub"].(string)
			if !ok {
				writeJSONError(w, APIError{Status: http.StatusUnauthorized, Code: "internal_error", Message: "Invalid user ID in token"})
				return
			}
			ctx := r.Context()
			ctx = context.WithValue(ctx, userIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			writeJSONError(w, APIError{Status: http.StatusUnauthorized, Code: "invalid_token", Message: "Invalid authorization token"})
			return
		}
	})
}

// loginHandler 处理登录请求，生成 JWT
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// 解析请求体（假设是 JSON 格式
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		writeJSONError(w, APIError{Status: http.StatusBadRequest, Code: "invalid_request", Message: "Invalid request body"})
		return
	}
	// 验证用户名和密码（这里使用硬编码的示例）
	expectedPassword, ok := users[credentials.Username]
	if !ok || credentials.Password != expectedPassword {
		writeJSONError(w, APIError{Status: http.StatusUnauthorized, Code: "invalid_credentials", Message: "Invalid username or password"})
		return
	}
	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": credentials.Username,                  // 使用 "sub" （subject）存储用户 ID
		"exp": time.Now().Add(time.Hour * 24).Unix(), // 过期时间为 24 小时
	})
	// 签名 JWT
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		writeJSONError(w, APIError{Status: http.StatusInternalServerError, Code: "internal_error", Message: "Failed to generate token"})
		return
	}
	// 返回 JWT
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

// protectedHandler 需要鉴权的受保护资源
func protectedHandler(w http.ResponseWriter, r *http.Request) {
	// 从 context 中获取用户 ID
	userID := r.Context().Value(userIDKey).(string)

	// 返回受保护的资源（这里只是示例）
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("Hello, %s! This is a protected resource.", userID)})
}

// homeHandler 示例：未受保护的公共资源
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // 标准的做法是明确设置 Content-Type
	w.WriteHeader(http.StatusOK)                       // 明确设置状态码是个好习惯
	response := map[string]string{
		"message": "Welcome to the homepage!",
	}

	// 使用 json.NewEncoder 进行编码，效率更高
	if err := json.NewEncoder(w).Encode(response); err != nil {
		// 极端情况下，如果 JSON 编码失败，应该记录错误并返回一个内部服务器错误
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {
	r := mux.NewRouter()

	// 注册中间件（注意顺序，先日志，后鉴权）
	r.Use(loggingMiddleware)

	// 未受保护的路由
	r.HandleFunc("/", homeHandler).Methods("GET")        // 主页，对“/”路径的 GET 请求做出响应
	r.HandleFunc("/login", loginHandler).Methods("POST") // 登录，对“/login”路径的 POST 请求做出响应

	// 受保护的路由
	r.Handle("/protected", authMiddleware(http.HandlerFunc(protectedHandler))).Methods("GET") // 受保护的资源，对“/protected”路径的 GET 请求做出响应

	// 启动服务器
	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
