package dao

import (
	"ds_test/geo/model"
	"ds_test/shadow-framework/orm/jinzhu/gorm"
	"errors"
)

type GeoTableColumnsDao struct {
	db *gorm.DB
}

func NewGeoTableColumnsDao(db *gorm.DB) *GeoTableColumnsDao {
	return &GeoTableColumnsDao{
		db: db,
	}
}

func (dao *GeoTableColumnsDao) Create(m *model.GeoTableColumns) error {
	return dao.db.Create(m).Error
}

func (dao *GeoTableColumnsDao) Find(m *model.GeoTableColumns) (result []model.GeoTableColumns, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *GeoTableColumnsDao) FindOne(m *model.GeoTableColumns) error {
	return dao.db.First(m, m).Error
}

func (dao *GeoTableColumnsDao) FindLast(m *model.GeoTableColumns) error {
	return dao.db.Last(m, m).Error
}

func (dao *GeoTableColumnsDao) FindPage(m *model.GeoTableColumns, rowbound model.RowBound) (result []model.GeoTableColumns, count int, err error) {
	err = dao.db.Model(&model.GeoTableColumns{}).Count(&count).Limit(rowbound.Limit).Offset(rowbound.Offset).Find(&result, m).Error
	return
}

func (dao *GeoTableColumnsDao) Get(m *model.GeoTableColumns) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *GeoTableColumnsDao) BatchGet(idbatch []int64) (result []model.GeoTableColumns, err error) {
	err = dao.db.Model(&model.GeoTableColumns{}).Where("ID in (?)", idbatch).Find(&result).Error
	return
}

func (dao *GeoTableColumnsDao) GetForUpdate(m *model.GeoTableColumns) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *GeoTableColumnsDao) Save(m *model.GeoTableColumns) error {
	return dao.db.Save(m).Error
}

func (dao *GeoTableColumnsDao) Delete(m *model.GeoTableColumns) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Unscoped().Delete(m).Error

}

func (dao *GeoTableColumnsDao) SoftDelete(m *model.GeoTableColumns) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *GeoTableColumnsDao) BatchDelete(idbatch []int64) error {
	return dao.db.Unscoped().Where("ID in (?)", idbatch).Delete(&model.GeoTableColumns{}).Error

}

func (dao *GeoTableColumnsDao) SoftBatchDelete(idbatch []int64) error {
	return dao.db.Where("ID in (?)", idbatch).Delete(&model.GeoTableColumns{}).Error
}

func (dao *GeoTableColumnsDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.GeoTableColumns{}).Where("ID = ?", id).Updates(attrs).Error
}

func (dao *GeoTableColumnsDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&model.GeoTableColumns{}).Where("ID = ?", id).Update(attr, value).Error
}

func (dao *GeoTableColumnsDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.GeoTableColumns{}).Where("ID in (?)", idbatch).Updates(attrs).Error
}

func (dao *GeoTableColumnsDao) Found(m *model.GeoTableColumns) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
