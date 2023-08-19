package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}
type OpeningResponse struct {
	ID        uint      `json:"id"`
	CreateAt  time.Time `json:"createAt"`
	DeletedAt time.Time `json:"deletedAt, omitempity"`
	UpdateAt  time.Time `json:"updateAt"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Remote    bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    int64     `json:"salary"`
}
