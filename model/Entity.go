package model

import (
	"time"
)

type GeoTable struct {
	ID              int64             `gorm:"primary_key" json:"id"`
	CreatedAt       *time.Time        `json:"createdAt"`
	UpdatedAt       *time.Time        `json:"updatedAt"`
	DeletedAt       *time.Time        `json:"deletedAt"`
	TableName       string            `gorm:"table_name"       json:"table_name"`      // 表名称
	TableComment    string            `gorm:"table_comment"    json:"table_comment"`   // 表描述
	ClassName       string            `gorm:"class_name"       json:"class_name"`      // 实体类名称
	TplCategory     string            `gorm:"tpl_category"     json:"tpl_category"`    // 使用的模板（crud单表操作 tree树表操作）
	PackageName     string            `gorm:"package_name"     json:"package_name"`    // 生成包路径
	ModuleName      string            `gorm:"module_name"      json:"module_name"`     // 生成模块名
	BusinessName    string            `gorm:"business_name"    json:"business_name"`   // 生成业务名
	FunctionName    string            `gorm:"function_name"    json:"function_name"`   // 生成功能名
	FunctionAuthor  string            `gorm:"function_author"  json:"function_author"` // 生成功能作者
	Options         string            `gorm:"options"          json:"options"`         // 其它生成选项
	CreateBy        string            `gorm:"create_by"        json:"create_by"`       // 创建者
	UpdateBy        string            `gorm:"update_by"        json:"update_by"`       // 更新者
	Remark          string            `gorm:"remark"           json:"remark"`          // 备注
	GeoTableColumns []GeoTableColumns `gorm:"FOREIGNKEY:TableID;ASSOCIATION_FOREIGNKEY:ID"`
}

type GeoTableColumns struct {
	ID            int64      `gorm:"primary_key" json:"id"`
	CreatedAt     *time.Time `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
	DeletedAt     *time.Time `json:"deletedAt"`
	TableID       int64      `gorm:"table_id"          json:"table_id"`       // 归属表编号
	ColumnName    string     `gorm:"column_name"       json:"column_name"`    // 列名称
	ColumnComment string     `gorm:"column_comment"    json:"column_comment"` // 列描述
	ColumnType    string     `gorm:"column_type"       json:"column_type"`    // 列类型
	GoType        string     `gorm:"go_type"           json:"go_type"`        // GO类型
	GoField       string     `gorm:"go_field"          json:"go_field"`       // GO字段名
	HtmlField     string     `gorm:"html_field"        json:"html_field"`     // html字段名
	IsPk          string     `gorm:"is_pk"             json:"is_pk"`          // 是否主键（1是）
	IsIncrement   string     `gorm:"is_increment"      json:"is_increment"`   // 是否自增（1是）
	IsRequired    string     `gorm:"is_required"       json:"is_required"`    // 是否必填（1是）
	IsInsert      string     `gorm:"is_insert"         json:"is_insert"`      // 是否为插入字段（1是）
	IsEdit        string     `gorm:"is_edit"           json:"is_edit"`        // 是否编辑字段（1是）
	IsList        string     `gorm:"is_list"           json:"is_list"`        // 是否列表字段（1是）
	IsQuery       string     `gorm:"is_query"          json:"is_query"`       // 是否查询字段（1是）
	QueryType     string     `gorm:"query_type"        json:"query_type"`     // 查询方式（等于、不等于、大于、小于、范围）
	HtmlType      string     `gorm:"html_type"         json:"html_type"`      // 显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）
	DictType      string     `gorm:"dict_type"         json:"dict_type"`      // 字典类型
	Sort          int        `gorm:"sort"              json:"sort"`           // 排序
	CreateBy      string     `gorm:"create_by"         json:"create_by"`      // 创建者
	UpdateBy      string     `gorm:"update_by"         json:"update_by"`      // 更新者
}
