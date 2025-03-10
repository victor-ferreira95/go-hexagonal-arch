package application

type ProductService struct {
	Persistence ProductPersistenceInterface
}

//  não precisa saber qual o objeto de persistencia que vai ser usado, ele vai usar a interface
func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
