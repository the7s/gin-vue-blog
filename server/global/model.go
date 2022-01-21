package global

import (
	"time"
)

type COMMON_MODEL struct {
	ID        uint `gorm:"primarykey"` //主键id
	CreatedAt MyTime
	UpdatedAt MyTime
	DeletedAt *time.Time `gorm:"index" json:"-"`
}
