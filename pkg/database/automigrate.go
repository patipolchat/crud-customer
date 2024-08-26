package database

import (
	"crud-customer/internal/entity"
	"github.com/go-faker/faker/v4"
	"math/rand"
)

func (g *gormDB) AutoMigrate() error {
	entityList := entity.GetEntityList()
	for _, e := range entityList {
		if err := g.db.AutoMigrate(e); err != nil {
			return err
		}
	}
	return nil
}

func (g *gormDB) Seed() error {
	for i := 0; i < 10; i++ {
		name := faker.Name()
		age := uint(rand.Intn(100) + 1)
		customer := entity.Customer{
			Name: &name,
			Age:  &age,
		}
		if err := g.db.Create(&customer).Error; err != nil {
			return err
		}
	}
	return nil
}
