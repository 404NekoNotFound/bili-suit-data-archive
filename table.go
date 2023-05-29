package table

import (
	"gorm.io/gorm"
	"time"
)

// Rank 排名列表
type Rank struct {
	gorm.Model
	ItemId int32     `json:"item_id" gorm:"index:,option:CONCURRENTLY"`
	Number int32     `json:"number" gorm:"index:,option:CONCURRENTLY"`
	Mid    uint64    `json:"mid" gorm:"index:,option:CONCURRENTLY"`
	Owner  string    `json:"owner" gorm:"index:,option:CONCURRENTLY"`
	Time   time.Time `json:"time"`
}

// Dynamic 动态
type Dynamic struct {
	gorm.Model
	ItemId       int32     `json:"item_id" gorm:"index:,option:CONCURRENTLY"`
	Number       int32     `json:"number" gorm:"index:,option:CONCURRENTLY"`
	Uid          uint64    `json:"uid"`
	Uname        string    `json:"uname"`
	DynamicStrId string    `json:"dynamic_str_id" gorm:"uniqueindex:,option:CONCURRENTLY"`
	BusinessId   string    `json:"business_id"`
	Time         time.Time `json:"time"`
}

// History 历史记录
type History struct {
	gorm.Model
	ItemId int32     `json:"item_id" gorm:"index:,option:CONCURRENTLY"`
	Number int32     `json:"number" gorm:"index:,option:CONCURRENTLY"`
	Mid    uint64    `json:"mid" gorm:"index:,option:CONCURRENTLY"`
	Owner  string    `json:"owner"`
	Time   time.Time `json:"time"`
}
