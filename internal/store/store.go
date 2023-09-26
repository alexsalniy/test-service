package store

type Store interface {
	ExtFIO() ExtendedFIORepository
}