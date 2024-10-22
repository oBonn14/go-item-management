package customer

type Customer struct {
	IDCustomer   int    `gorm:"column:id_customer;primaryKey;autoIncrement" json:"id_customer" `
	CustomerName string `gorm:"column:customer_name;type:varchar(50)" json:"name" `
	Age          int    `gorm:"column:age;type:int" json:"age" `
}

func (Customer) TableName() string {
	return "customer"
}

type Customers []Customer
