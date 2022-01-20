package global

import (
	"time"
)

type COMMON_MODEL struct {
	ID        uint `gorm:"primarykey"` //主键id
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
