package persistentobject

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:char(36);not null;uniqueIndex;primaryKey"`
	Username  string    `gorm:"type:varchar(100);not null;uniqueIndex"`
	Email     string    `gorm:"type:varchar(100);not null;uniqueIndex"`
	Password  string    `gorm:"type:varchar(255);not null"`
	IsActive  bool      `gorm:"default:true;type:bool"`
	Roles     []Role    `gorm:"many2many:user_roles"` // many-to-many relationship
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` // Soft delete
}

// TableName overrides the table name
func (u *User) TableName() string {
	return "go_db_user"
}

// BeforeCreate hook will auto-generate UUID if it's not set
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}
