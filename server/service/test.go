package service

import (
	"server/model"
)

type TestService struct{}

func (TestService) SetTest(book *model.Test) error {
	_, err := DbEngine.Insert(book)
	if err != nil {
		return err
	}
	return nil
}

func (TestService) GetTestList() []model.Test {
	tests := make([]model.Test, 0)
	err := DbEngine.Distinct("id", "name").Limit(10, 0).Find(&tests)
	if err != nil {
		panic(err)
	}
	return tests
}
