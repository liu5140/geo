// ==========================================================================
// 生成日期：2020-03-12 18:05:41 &#43;0800 CST
// 生成路径: app/model/model/Answer/Answer_entity.go
// 生成人：yunjie
// ==========================================================================

package dao

import (
	"errors"
	"yj-app/data/model"
)

func (dao *AnswerDao) SearchAnswers(condition *model.AnswerSearchCondition, rowbound *model.RowBound) (result []model.Answer, count int, err error) {
	db := dao.db
	if len(condition.UserIDs) > 0 {
		db = db.Where("id in (?) ", condition.UserIDs)
	}

	if rowbound == nil {
		err = db.Model(&model.Answer{}).Order("ID desc").Count(&count).Find(&result).Error
	} else {
		err = db.Model(&model.Answer{}).Order("ID desc").Count(&count).Offset(rowbound.Offset).Limit(rowbound.Limit).Find(&result).Error
	}

	return result, count, err

}