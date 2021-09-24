package service

import (
	"server/config"
	"server/model"
)

type TestService struct{}

var test model.Test
var tests []model.Test

func (TestService) SetTest(test *model.Test) error {
	db := config.GetConn()
	defer db.Close()
	result := db.Create(&test)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (TestService) GetTestList() []model.Test {
	db := config.GetConn()
	defer db.Close()
	db.Find(&tests)
	return tests
}

func (TestService) UpdateTest(newTest *model.Test) error {
	db := config.GetConn()
	defer db.Close()
	db.First(&test)
	result := db.Save(newTest)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (TestService) DeleteTest(id int) error {
	db := config.GetConn()
	defer db.Close()
	db.First(&test)
	result := db.Delete(&test, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
