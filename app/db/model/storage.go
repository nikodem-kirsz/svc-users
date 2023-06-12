package model

type Storage interface {
}

type MysqlStorage struct {
}

func NewMySqlStorage() Storage {
	return &MysqlStorage{}
}
