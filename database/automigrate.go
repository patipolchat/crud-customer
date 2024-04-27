package database

import "crud-customer/entity"

func (g *gormDB) AutoMigrate() error {
	entityList := entity.GetEntityList()
	for _, e := range entityList {
		if err := g.db.AutoMigrate(e); err != nil {
			return err
		}
	}
	return nil
}
