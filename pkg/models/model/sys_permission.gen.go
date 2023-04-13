// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSysPermission = "sys_permission"

// SysPermission mapped from table <sys_permission>
type SysPermission struct {
	ID        int64          `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Name      string         `gorm:"column:name" json:"name"`
	State     int16          `gorm:"column:state" json:"state"`
	Remark    string         `gorm:"column:remark" json:"remark"`
	Type      int16          `gorm:"column:type" json:"type"` // 1：菜单，2：按钮不能为空
}

// TableName SysPermission's table name
func (*SysPermission) TableName() string {
	return TableNameSysPermission
}
