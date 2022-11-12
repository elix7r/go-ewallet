package transaction

type Storage interface {
	GetAll() ([]*Transaction, error)
	GetByIssued() ([]*Transaction, error)
	GetByUUID(uuid string) (*Transaction, error)
	Create(from string, to string, amount float64, tType Type) (*Transaction, error)
	UpdateIsIssuedByUUID(uuid string) error
	DeleteByUUID(uuid string) (string, error)
}
