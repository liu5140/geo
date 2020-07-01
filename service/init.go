package service

import (
	"html/template"
	"sync"

	"ds_test/geo/model"
	"ds_test/shadow-framework/logger"
	"ds_test/shadow-framework/orm/jinzhu/gorm"

	"github.com/leekchan/accounting"
	"github.com/shopspring/decimal"
)

var (
	Log *logger.Logger
	l   sync.Mutex
)
var TempFuncMap template.FuncMap

func init() {
	Log = logger.InitLog()
	TempFuncMap = make(template.FuncMap)
	TempFuncMap["CheckField"] = CheckField
}

const (
	TIME_FORMAT_WITH_MS         = "2006-01-02 15:04:05.000"
	TIME_FORMAT                 = "2006-01-02 15:04:05"
	TIME_FORMAT_WO_SEC_COMPACT  = "200601021504"
	TIME_FORMAT_COMPACT         = "20060102150405"
	TIME_FORMAT_WITH_MS_COMPACT = "20060102150405000"
	DATE_FORMAT                 = "2006-01-02"
	DATE_FORMAT_COMPACT         = "20060102"
	MONTH_FORMAT                = "2006-01"
)

const (
	HSET         = "HSET"
	HINCRBYFLOAT = "HINCRBYFLOAT"
	INCRBYFLOAT  = "INCRBYFLOAT"
	HINCRBY      = "HINCRBY"
	SADD         = "SADD"
	KEYS         = "KEYS"
	HGETALL      = "HGETALL"
	SISMEMBER    = "SISMEMBER"
	HMSET        = "HMSET"
	HMGET        = "HMGET"
	GET          = "GET"
)

func CheckField(fields []model.GeoTableColumns, fieldName string) bool {
	for _, f := range fields {
		if f.GoField == fieldName {
			return true
		}
	}
	return false
}

func closeTx(tx *gorm.DB, err *error) {
	r := recover()
	if r != nil {
		tx.Rollback()
		Log.Error(r)
		return
	}

	if *err != nil {
		tx.Rollback()
		Log.Errorf("%+v", *err)
		return
	}
	tx.Commit()
}

func CloseTx(tx *gorm.DB, err *error) {
	closeTx(tx, err)
}

type TanslateModel struct {
	Number int
	Value  string
}

func MoneyFormatDecimal(money decimal.Decimal, precision int) string {
	ac := accounting.Accounting{Symbol: "", Precision: precision}
	m, _ := money.Float64()
	return ac.FormatMoney(m)
}

func MoneyFormatFloat(money float64, precision int) string {
	ac := accounting.Accounting{Symbol: "", Precision: precision}
	return ac.FormatMoney(money)
}
