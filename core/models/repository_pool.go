package models

import "time"

// 中心存储池中的文件信息
type RepositoryPool struct {
	Id        int
	Identity  string
	Hash      string
	Name      string
	Ext       string
	Size      int64
	Path      string
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

// 将结构体对应数据库中对应的表名称
func (table RepositoryPool) TableName() string {
	return "repository_pool"
}
