package data

import (
	"github.com/John-Dembaremba/KAFKA-PRACTICE.git/schema"
	fakerInstance "github.com/jaswdr/faker/v2"
)

type FakeData interface {
	getPersonData() *schema.Person
	getCompanyData() *schema.Company
}

type Handler struct {
	fake *fakerInstance.Faker
}

func New() *Handler {
	fake := fakerInstance.New()
	return &Handler{fake: &fake}
}

func (f *Handler) getPersonData() *schema.Person {
	return &schema.Person{
		ID:      f.fake.Int(),
		Name:    f.fake.Person().FirstName(),
		Surname: f.fake.Person().LastName(),
		Email:   f.fake.Person().Contact().Email,
	}
}

func (f *Handler) getCompanyData() *schema.Company {
	return &schema.Company{
		ID:   f.fake.Int(),
		Name: f.fake.Company().Name(),
	}
}
