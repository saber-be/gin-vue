package databases

type IDBAdapter interface {
	Connect() error
	Close() error
	Create(model interface{}) error
	Find(models interface{}, conds ...interface{}) error
	Update(model interface{}, values interface{}) error
	Delete(model interface{}) error
}
