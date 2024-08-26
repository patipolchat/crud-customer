package entity

type Customer struct {
	ID   uint    `json:"id" gorm:"primaryKey;autoIncrement;not null;index"`
	Name *string `json:"name" gorm:"not null"`
	Age  *uint   `json:"age" gorm:"not null"`
}

func init() {
	entityList = append(entityList, Customer{})
}
