package dao

import (
	"ds_test/geo/model"
)

func (dao *GeoTableDao) GetGeoTableByID(m *model.GeoTable) error {
	return dao.db.Model(&model.GeoTable{}).Where("id = ?", m.ID).Preload("GeoTableColumns").Find(&m).Error
}

func (dao *GeoTableDao) SelectDbTableListByNames(tableNames []string) (result []model.GeoTable, err error) {
	db := dao.db
	db.Where("table_name NOT LIKE 'qrtz_%'")
	db.Where("table_name NOT LIKE 'gen_%'")
	db.Where("table_schema = (select database())")
	if len(tableNames) > 0 {
		db.Where("table_name in (?)", tableNames)
	}
	if err = db.Table(" information_schema.tables ").Find(&result).Error; err != nil {
		return []model.GeoTable{}, err
	}
	return result, nil
}

func (dao *GeoTableColumnsDao) SelectDbTableColumnsByName(tableName string) (result []model.GeoTableColumns, err error) {
	db := dao.db
	err = db.Raw("select column_name,"+
		"(case when (is_nullable = 'no' && column_key != 'PRI') then '1' else null end) as is_required,"+
		"(case when column_key = 'PRI' then '1' else '0' end) as is_pk,"+
		"ordinal_position as sort,"+
		"column_comment,"+
		"(case when extra = 'auto_increment' then '1' else '0' end) as is_increment,"+
		"column_type "+
		"from information_schema.columns where table_schema = (select database()) and table_name= ? order by ordinal_position", tableName).Scan(&result).Error
	if err != nil {
		return []model.GeoTableColumns{}, err
	}
	return result, nil
}
