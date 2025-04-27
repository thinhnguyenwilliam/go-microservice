package persistentobject

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gorm:"primaryKey;autoIncrement;not null;comment:Role ID Is Here"`
	RoleName string `gorm:"type:varchar(100);not null;uniqueIndex"`
	RoleNote string `gorm:"type:varchar(255);not null"`
}

// TableName overrides the table name for Role
func (r *Role) TableName() string {
	return "go_db_role"
}
