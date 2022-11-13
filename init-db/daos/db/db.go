package db

type DAO struct {
}

type DAOInterface interface {
	CheckDBs() (exists bool, err error)
	DeleteDBs() (err error)
	CreateDBs() (err error)
}

func NewDAO() DAOInterface {
	return &DAO{}
}

func (d *DAO) CheckDBs() (exists bool, err error) {
	// TODO
	return
}

func (d *DAO) DeleteDBs() (err error) {
	// TODO
	return
}

func (d *DAO) CreateDBs() (err error) {
	// TODO
	return
}
