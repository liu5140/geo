package main

import (
	"ds_test/geo/model"
	"ds_test/geo/service"
	"ds_test/shadow-framework/logger"
	"ds_test/shadow-framework/orm/datasource"
	_ "ds_test/shadow-framework/orm/datasource/datasourcemanager"
	_ "ds_test/shadow-framework/orm/jinzhu/gorm/dialects/mysql"
	"sync"
)

var (
	wg           sync.WaitGroup
	BuildVersion string
	Log          *logger.Logger = logger.InitLog()
)

func main() {

	d := []interface{}{
		new(model.GeoTable),
		new(model.GeoTableColumns),
	}
	datasource.DatasourceManagerInstance(datasource.DATASOURCE_MANAGER).RegisterModels(d...)

	config := service.NewTableService()
	config.GenCode(53)
}
