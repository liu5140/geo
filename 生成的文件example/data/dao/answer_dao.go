// ==========================================================================
// 生成日期：2020-03-12 18:05:41 &#43;0800 CST
// 生成路径: app/model/model/Answer/Answer_entity.go
// 生成人：yunjie
// ==========================================================================

package dao

import (
	"errors"
	"yj-app/data/model"
	"github.com/shopspring/decimal"
	"yj-app/shadow-framework/orm/jinzhu/gorm"
)

type AnswerDao struct {
	db *gorm.DB
}

func NewAnswerDao(db *gorm.DB) *AnswerDao {
	return &AnswerDao{
		db: db,
	}
}

func (dao *AnswerDao) Create(m *model.Answer) error {
	return dao.db.Create(m).Error
}

func (dao *AnswerDao) Find(m *model.Answer) (result []model.Answer, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *AnswerDao) FindOne(m *model.Answer) error {
	return dao.db.First(m, m).Error
}

func (dao *AnswerDao) FindLast(m *model.Answer) error {
	return dao.db.Last(m, m).Error
}

func (dao *AnswerDao) FindPage(m *model.Answer, rowbound model.RowBound, desc bool)  (result []model.Answer, count int, err error)  {
	db := dao.db
	if desc {
		db = db.Order("id desc")
	}
	err = db.Model(m).Count(&count).Limit(rowbound.Limit).Offset(rowbound.Offset).Find(&result, m).Error
	return
}

func (dao *AnswerDao) Get(m *model.Answer) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *AnswerDao) BatchGet(idbatch []int64) (result []model.Answer, err error) {
	if len(idbatch) == 0 {
		return nil, errors.New("id is nil")
	}
	err = dao.db.Model(&model.Answer{}).Where("ID in (?)", idbatch).Find(&result).Error
	return
}

func (dao *AnswerDao) GetForUpdate(m *model.Answer) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *AnswerDao) Save(m *model.Answer) error {
	return dao.db.Save(m).Error
}

func (dao *AnswerDao) Delete(m *model.Answer) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}



func (dao *AnswerDao) BatchDelete(idbatch []int64) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Where("ID in (?)", idbatch).Delete(&model.Answer{}).Error
}



func (dao *AnswerDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.Answer{}).Where("ID = ?", id).Updates(attrs).Error
}

func (dao *AnswerDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&model.Answer{}).Where("ID = ?", id).Update(attr, value).Error
}

func (dao *AnswerDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Model(&model.Answer{}).Where("ID in (?)", idbatch).Updates(attrs).Error
}

func (dao *AnswerDao) Found(m *model.Answer) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}