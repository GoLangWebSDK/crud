package crud

type Repository[T any] interface {
	All() ([]T, error)
	Create(model T) error
	Read(ID uint32) (T, error)
	Update(ID uint32, model T) error
	Delete(ID uint32) error
}
