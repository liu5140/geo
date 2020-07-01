package service

import (
	"ds_test/app/data/dao"
	"ds_test/shadow-framework/orm/datasource"
	"yj-app/data/model"
)

type AnswerService struct {
}

var answerService *AnswerService

func NewAnswerService() *AnswerService {
	if answerService == nil {
		l.Lock()
		answerService = &AnswerService{}
		l.Unlock()
	}
	return answerService
}

func (service AnswerService) SearchAnswerPaging(condition *model.AnswerSearchCondition, pageNum int, pageSize int) (request []model.Answer, count int, err error) {
	rowbound := model.NewRowBound(pageNum, pageSize)
	return service.searchAnswer(condition, &rowbound)
}

func (service AnswerService) SearchAnswerOutPaging(condition *model.AnswerSearchCondition) (request []model.Answer, count int, err error) {
	return service.searchAnswer(condition, nil)
}

func (service AnswerService) searchAnswer(condition *model.AnswerSearchCondition, rowbound *model.RowBound) (request []model.Answer, count int, err error) {
	db := datasource.DatasourceManagerInstance(datasource.DATASOURCE_MANAGER).Datasource()
	messageDao := dao.NewAnswerDao(db)
	request, count, err = messageDao.SeachAnswer(condition, rowbound)
	if err != nil {
		Log.Error(err)
		return nil, 0, err
	}
	return request, count, nil
}

func (service AnswerService) DeleteAnswer(id int64) (err error) {
	db := datasource.DatasourceManagerInstance(datasource.DATASOURCE_MANAGER).Datasource()
	messageDao := dao.NewAnswerDao(db)
	return messageDao.Delete(&model.Answer{ID: id})
}

func (service AnswerService) CreateAnswer(m *model.Answer) (err error) {
	db := datasource.DatasourceManagerInstance(datasource.DATASOURCE_MANAGER).Datasource()
	messageDao := dao.NewAnswerDao(db)
	return messageDao.Create(m)
}
