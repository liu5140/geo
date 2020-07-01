package service

import (
	"ds_test/geo/dao"
	"ds_test/geo/model"
	"ds_test/shadow-framework/orm/datasource"
	"errors"
	"html/template"
	"os"
	"strings"

	"github.com/gogf/gf/util/gconv"
)

type TableService struct {
}

var tableService *TableService

func NewTableService() *TableService {
	if tableService == nil {
		l.Lock()
		if tableService == nil {
			tableService = &TableService{}
		}
		l.Unlock()
	}
	return tableService
}

func (service TableService) ImportGenTable(tableList []model.GeoTable, operName string) (err error) {
	db := datasource.DatasourceManagerInstance(datasource.DATASOURCE_MANAGER).Datasource()
	tx := db.Begin()
	defer closeTx(tx, &err)
	enDao := dao.NewGeoTableDao(tx)
	encDao := dao.NewGeoTableColumnsDao(tx)
	if tableList != nil && operName != "" {
		for _, table := range tableList {
			tableName := table.TableName
			err := enDao.Create(&table)
			if err != nil {
				return err
			}
			// 保存列信息
			genTableColumns, err := encDao.SelectDbTableColumnsByName(tableName)

			if err != nil || len(genTableColumns) <= 0 {
				return errors.New("获取列数据失败")
			}

			for _, column := range genTableColumns {
				column.TableID = table.ID
				column.CreateBy = table.CreateBy
				err = encDao.Create(&column)
				if err != nil {
					return errors.New("获取列数据失败")
				}
			}
		}
	} else {
		return errors.New("参数错误")
	}
	return nil
}

func (service TableService) GenCode(tableid int64) (err error) {
	db := datasource.DatasourceManagerInstance(datasource.DATASOURCE_MANAGER).Datasource()
	enDao := dao.NewGeoTableDao(db)
	entityTable := model.GeoTable{
		ID: tableid,
	}
	err = enDao.GetGeoTableByID(&entityTable)
	if err != nil {
		return err
	}
	//生成model 文件
	err = service.genModel(entityTable)
	if err != nil {
		return err
	}
	//生成Dao 文件
	err = service.genDao(entityTable)
	if err != nil {
		return err
	}
	//生成DaoExt 文件
	err = service.genDaoExt(entityTable)
	if err != nil {
		return err
	}
	//生成Dto 文件
	err = service.genDto(entityTable)
	if err != nil {
		return err
	}

	//生成Service 文件
	err = service.genService(entityTable)
	if err != nil {
		return err
	}

	return nil
}

//genModel 生成model文件
func (service TableService) genModel(entityTable model.GeoTable) (err error) {
	tmpl, err := template.New("model.html").Funcs(TempFuncMap).ParseFiles("geo/template/model.html")
	if err != nil {
		Log.Error(err)
		return err
	}
	//获取当前运行时目录
	curDir, err := os.Getwd()
	if err != nil {
		return err
	}
	fileName := strings.Join([]string{curDir, "/app/data/model/", entityTable.ClassName, "_entity.go"}, "")
	return service.genExecute(tmpl, entityTable, fileName)
}

//service 生成service文件
func (service TableService) genService(entityTable model.GeoTable) (err error) {
	tmpl, err := template.New("service.html").Funcs(TempFuncMap).ParseFiles("geo/template/service.html")
	if err != nil {
		Log.Error(err)
		return err
	}
	//获取当前运行时目录
	curDir, err := os.Getwd()
	if err != nil {
		return err
	}
	fileName := strings.Join([]string{curDir, "/app/service/", entityTable.ClassName, "_service.go"}, "")
	return service.genExecute(tmpl, entityTable, fileName)
}

//dto 生成dto文件
func (service TableService) genDto(entityTable model.GeoTable) (err error) {
	tmpl, err := template.New("dto.html").Funcs(TempFuncMap).ParseFiles("geo/template/dto.html")
	if err != nil {
		Log.Error(err)
		return err
	}
	//获取当前运行时目录
	curDir, err := os.Getwd()
	if err != nil {
		return err
	}
	fileName := strings.Join([]string{curDir, "/app/service/dto/", entityTable.ClassName, "_dto.go"}, "")
	return service.genExecute(tmpl, entityTable, fileName)
}

//genDao 生成Dao文件
func (service TableService) genDao(entityTable model.GeoTable) (err error) {
	tmpl, err := template.New("dao.html").Funcs(TempFuncMap).ParseFiles("geo/template/dao.html")
	if err != nil {
		Log.Error(err)
		return err
	}
	//获取当前运行时目录
	curDir, err := os.Getwd()
	if err != nil {
		return err
	}
	fileName := strings.Join([]string{curDir, "/app/data/dao/", entityTable.ClassName, "_dao.go"}, "")
	return service.genExecute(tmpl, entityTable, fileName)
}

//genDao 生成Dao_ext文件
func (service TableService) genDaoExt(entityTable model.GeoTable) (err error) {
	tmpl, err := template.New("dao_ext.html").Funcs(TempFuncMap).ParseFiles("geo/template/dao_ext.html")
	if err != nil {
		Log.Error(err)
		return err
	}
	//获取当前运行时目录
	curDir, err := os.Getwd()
	if err != nil {
		return err
	}
	fileName := strings.Join([]string{curDir, "/app/data/dao/", entityTable.ClassName, "_dao_ext.go"}, "")
	return service.genExecute(tmpl, entityTable, fileName)
}

func (service TableService) genExecute(tmpl *template.Template, entityTable model.GeoTable, fileName string) (err error) {
	createFile(fileName)
	file, err := os.OpenFile(fileName, os.O_RDWR, os.ModeAppend)
	if err != nil {
		Log.Error(err)
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			Log.Error(err)
		}
	}()
	err = tmpl.Execute(file, map[string]model.GeoTable{"table": entityTable})
	if err != nil {
		Log.Error(err)
		return err
	}
	Log.Infoln("=============生成:", entityTable.BusinessName+"===========路径:", fileName)
	return nil
}

func createFile(fileName string) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		os.MkdirAll(fileName, 0744)
	}

	if _, err := os.Stat(fileName); err == nil {
		os.Remove(fileName)
	}
	_, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
}

//初始化列属性字段
func InitColumnField(column *model.GeoTableColumns) {
	dataType := GetDbType(column.ColumnType)
	columnName := column.ColumnName
	// column.TableId = table.TableId
	// column.CreateBy = table.CreateBy
	//设置字段名
	column.GoField = ConvertToCamelCase(columnName)
	column.HtmlField = ConvertToCamelCase1(columnName)

	if IsStringObject(dataType) {
		//字段为字符串类型
		column.GoType = "string"
		if strings.EqualFold(dataType, "text") || strings.EqualFold(dataType, "tinytext") || strings.EqualFold(dataType, "mediumtext") || strings.EqualFold(dataType, "longtext") {
			column.HtmlType = "textarea"
		} else {
			columnLength := GetColumnLength(column.ColumnType)
			if columnLength >= 500 {
				column.HtmlType = "textarea"
			} else {
				column.HtmlType = "input"
			}
		}
	} else if IsTimeObject(dataType) {
		//字段为时间类型
		column.GoType = "Time"
		column.HtmlType = "datatime"
	} else if IsNumberObject(dataType) {
		//字段为数字类型
		column.HtmlType = "input"
		// 如果是浮点型
		tmp := column.ColumnType
		if tmp == "float" || tmp == "double" {
			column.GoType = "float64"
		} else {
			start := strings.Index(tmp, "(")
			end := strings.Index(tmp, ")")
			result := tmp[start+1 : end]
			arr := strings.Split(result, ",")
			if len(arr) == 2 && gconv.Int(arr[1]) > 0 {
				column.GoType = "float64"
			} else if len(arr) == 1 && gconv.Int(arr[0]) <= 10 {
				column.GoType = "int"
			} else {
				column.GoType = "int64"
			}
		}
	}
	//新增字段
	if columnName == "create_by" || columnName == "create_time" || columnName == "update_by" || columnName == "update_time" {
		column.IsRequired = "0"
		column.IsInsert = "0"
	} else {
		column.IsRequired = "0"
		column.IsInsert = "1"
		if strings.Index(columnName, "name") >= 0 || strings.Index(columnName, "status") >= 0 {
			column.IsRequired = "1"
		}
	}

	// 编辑字段
	if IsNotEdit(columnName) {
		if column.IsPk == "1" {
			column.IsEdit = "0"
		} else {
			column.IsEdit = "1"
		}
	} else {
		column.IsEdit = "0"
	}
	// 列表字段
	if IsNotList(columnName) {
		column.IsList = "1"
	} else {
		column.IsList = "0"
	}
	// 查询字段
	if IsNotQuery(columnName) {
		column.IsQuery = "1"
	} else {
		column.IsQuery = "0"
	}

	// 查询字段类型
	if CheckNameColumn(columnName) {
		column.QueryType = "LIKE"
	} else {
		column.QueryType = "EQ"
	}

	// 状态字段设置单选框
	if CheckStatusColumn(columnName) {
		column.HtmlType = "radio"
	} else if CheckTypeColumn(columnName) || CheckSexColumn(columnName) {
		// 类型&性别字段设置下拉框
		column.HtmlType = "select"
	}
}
