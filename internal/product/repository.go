package product

// Repository abstracts data storage
type Repository interface {
	GetByID(id string) (*Product, error)
	ListAll() ([]*Product, error)
}

// InMemoryRepository is a simple in-memory store
type InMemoryRepository struct {
	data map[string]*Product
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		data: map[string]*Product{
			"1": {ID: "1", Name: "Laptop", Price: 1200},
			"2": {ID: "2", Name: "Phone", Price: 800},
		},
	}
}

func (r *InMemoryRepository) GetByID(id string) (*Product, error) {
	if p, ok := r.data[id]; ok {
		return p, nil
	}
	return nil, nil
}

func (r *InMemoryRepository) ListAll() ([]*Product, error) {
	products := []*Product{}
	for _, p := range r.data {
		products = append(products, p)
	}
	return products, nil
}
