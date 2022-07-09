package databases

import "qgen/challange/models"

type IDBAdapter interface {
	Connect() error
	Close() error
	Create(model interface{}) error
	Find(models interface{}, pagination *models.Pagination, conds ...interface{}) error
	Update(model interface{}, values interface{}) error
	Delete(model interface{}) error
}
