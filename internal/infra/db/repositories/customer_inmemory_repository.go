package repositories

import "github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/domain/entities"

type CustomerInMemoryRepository struct {
}

var (
	entitiesMap = make(map[int]entities.Customer)
)

func NewCustomerInMemoryRepository() *CustomerInMemoryRepository {
	entitiesMap[1] = *entities.NewCustomer(1, 0, 100000, []entities.Transaction{})
	return &CustomerInMemoryRepository{}
}
func (r *CustomerInMemoryRepository) FindById(id int) (entities.Customer, error) {
	return entitiesMap[id], nil
}
func (r *CustomerInMemoryRepository) Save(c entities.Customer) error {
	entitiesMap[c.ID()] = c
	return nil
}
