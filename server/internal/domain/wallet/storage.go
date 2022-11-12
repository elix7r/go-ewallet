package wallet

type Storage interface {
	GetAll() []*Wallet
	GetByUUID(uuid string) (*Wallet, error)
	Create() (*Wallet, error)
	Update(w *Wallet) (*Wallet, error)
	DeleteByUUID(uuid string) (string, error)
}
