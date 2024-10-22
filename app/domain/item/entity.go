package item

type Item struct {
	IdItem   string `gorm:"column:id_item;primary_key;autoIncrement;" json:"idItem"`
	CodeItem string `gorm:"column:code_item;type:varchar(50)" json:"codeItem"`
	NameItem string `gorm:"column:name_item;type:varchar(50)" json:"nameItem"`
	Qty      int    `gorm:"column:qty;type:int" json:"qty"`
	Price    int    `gorm:"column:price;type:int" json:"price"`
}

func (Item) TableName() string {
	return "item"
}

type Items []Item
