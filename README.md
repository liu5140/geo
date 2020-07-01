# geo
go语言根据数据库表结构,反向生成dao,services,router,model 等文件



记报错：

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


 template.New("model.html") 这个new里面的东西，要和ParseFiles("geo/template/model.html") 后面的一致才好，不然会报错误（template: "defaultModel" is an incomplete or empty template）
 
 
