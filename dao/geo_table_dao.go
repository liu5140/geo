package dao

import (
	"errors"

	"ds_test/geo/model"

	"ds_test/shadow-framework/orm/jinzhu/gorm"
)

type GeoTableDao struct {
	db *gorm.DB
}

func NewGeoTableDao(db *gorm.DB) *GeoTableDao {
	return &GeoTableDao{
		db: db,
	}
}

func (dao *GeoTableDao) Create(m *model.GeoTable) error {
	return dao.db.Create(m).Error
}

func (dao *GeoTableDao) Find(m *model.GeoTable) (result []model.GeoTable, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *GeoTableDao) FindOne(m *model.GeoTable) error {
	return dao.db.First(m, m).Error
}

func (dao *GeoTableDao) FindLast(m *model.GeoTable) error {
	return dao.db.Last(m, m).Error
}

func (dao *GeoTableDao) FindPage(m *model.GeoTable, rowbound model.RowBound) (result []model.GeoTable, count int, err error) {
	err = dao.db.Model(&model.GeoTable{}).Count(&count).Limit(rowbound.Limit).Offset(rowbound.Offset).Find(&result, m).Error
	return
}

func (dao *GeoTableDao) Get(m *model.GeoTable) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *GeoTableDao) BatchGet(idbatch []int64) (result []model.GeoTable, err error) {
	err = dao.db.Model(&model.GeoTable{}).Where("ID in (?)", idbatch).Find(&result).Error
	return
}

func (dao *GeoTableDao) GetForUpdate(m *model.GeoTable) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *GeoTableDao) Save(m *model.GeoTable) error {
	return dao.db.Save(m).Error
}

func (dao *GeoTableDao) Delete(m *model.GeoTable) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Unscoped().Delete(m).Error

}

func (dao *GeoTableDao) SoftDelete(m *model.GeoTable) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *GeoTableDao) BatchDelete(idbatch []int64) error {
	return dao.db.Unscoped().Where("ID in (?)", idbatch).Delete(&model.GeoTable{}).Error

}

func (dao *GeoTableDao) SoftBatchDelete(idbatch []int64) error {
	return dao.db.Where("ID in (?)", idbatch).Delete(&model.GeoTable{}).Error
}

func (dao *GeoTableDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.GeoTable{}).Where("ID = ?", id).Updates(attrs).Error
}

func (dao *GeoTableDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&model.GeoTable{}).Where("ID = ?", id).Update(attr, value).Error
}

func (dao *GeoTableDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.GeoTable{}).Where("ID in (?)", idbatch).Updates(attrs).Error
}

func (dao *GeoTableDao) Found(m *model.GeoTable) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
