package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Diego", "Giron", "68992", "2000-01-01", "1"},
		{"2002", "Rob", "Giron", "68992", "2000-01-01", "1"},
	}
	return CustomerRepositoryStub{customers}
}
