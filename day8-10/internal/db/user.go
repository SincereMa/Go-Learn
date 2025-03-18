package db

import (
	"database/sql"
	"myproject/pkg/models"
)

// UserRepository 实现 pkc/models 中定义的 UserRepository 接口
type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	// 定义 User 结构体变量
	user := &models.User{}
	// 执行数据库查询操作
	row := r.db.QueryRow("SELECT id, username FROM users WHERE id = ?", id)
	// 将查询结果映射到 User 结构体变量
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}
	return user, nil
}
