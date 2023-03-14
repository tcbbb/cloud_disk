package models

import "time"

// 中心存储池中的文件与特定用户进行关联
type UserRepository struct {
	Id                 int
	Identity           string // 用户关联后的唯一标识
	UserIdentity       string // 用户唯一标识
	ParentId           int64  // 父级目录
	RepositoryIdentity string // 中心存储池中文件唯一标识
	Ext                string
	Name               string
	CreatedAt          time.Time `xorm:"created"`
	UpdatedAt          time.Time `xorm:"updated"`
	DeletedAt          time.Time `xorm:"deleted"`
}

func (table UserRepository) TableName() string {
	return "user_repository"
}
