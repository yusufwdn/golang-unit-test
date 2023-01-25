package repository

import (
	"golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

// sangat disarankan untuk melakukan unit testing pada fungsi agak sulit seperti koneksi ke pihak ke 3
func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	// parameter yang dimasukkan di Called() disesuaikan dengan fungsi yang memanggilnya.
	// contoh: FindById(id, id2) berarti Mock.Called(id, id2)
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(entity.Category)
		return &category
	}
}
