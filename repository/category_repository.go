package repository

import "golang-unit-test/entity"

// membuat kontrak
type CateoryRepository interface {
	FindById(id string) *entity.Category
}
