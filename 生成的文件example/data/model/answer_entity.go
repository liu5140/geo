// ==========================================================================
// 生成日期：2020-03-12 18:05:41 &#43;0800 CST
// 生成路径: app/model/model/Answer/Answer_entity.go
// 生成人：yunjie
// ==========================================================================

package model

import (
	"time"
)

// Answer is the golang structure for table t_answer.
type Answer struct {
	// 主键
	Id int64 `gorm:"id,primary" json:"id"`
	// 问题ID
	Pid int64 `gorm:"pid" json:"pid"`
	// 回复人类别
	Atype int `gorm:"atype" json:"atype"`
	// 回复人ID
	UserId int64 `gorm:"user_id" json:"user_id"`
	// 回复人名称
	NickName string `gorm:"nick_name" json:"nick_name"`
	// 回复人头像
	Avatar string `gorm:"avatar" json:"avatar"`
	// 回复内容
	Remark string `gorm:"remark" json:"remark"`
	// 回复图片1
	Img1 string `gorm:"img1" json:"img1"`
	// 回复图片2
	Img2 string `gorm:"img2" json:"img2"`
	// 回复图片3
	Img3 string `gorm:"img3" json:"img3"`
	// 状态
	Status int `gorm:"status" json:"status"`
	// 创建时间
	CreateTime *time.Time `gorm:"created_at" json:"created_at"`
	// 更新时间
	UpdateTime *time.Time `gorm:"updated_at" json:"updated_at"`
}

func (model Answer) TableName() string {
	return "t_answer"
}

type AnswerSearchCondition struct {
	Pid      int64  // 问题ID
	Atype    int    // 回复人类别
	UserId   int64  // 回复人ID
	NickName string // 回复人名称
	Avatar   string // 回复人头像
	Remark   string // 回复内容
	Img1     string // 回复图片1
	Img2     string // 回复图片2
	Img3     string // 回复图片3
	Status   int    // 状态

}
