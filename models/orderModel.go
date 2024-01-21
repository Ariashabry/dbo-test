package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID          uint          `json:"id" gorm:"column:id_order;primaryKey;autoIncrement;"`
	CustomerID  uint          `json:"customerID" gorm:"column:id_user;foreignKey:ID"`
	Cart        []HistoryCart `json:"cart" gorm:"foreignKey:IDOrder"`
	CreatedAt   string        `json:"created_at" gorm:"column:created_at;type:datetime"`
	UpdatedAt   string        `json:"updated_at" gorm:"column:updated_at;type:datetime"`
	TotalAmount float64       `json:"totalAmount" gorm:"column:total_amount"`
}

type Product struct {
	ID          uint    `json:"id" gorm:"column:id_product;primaryKey;autoIncrement;"`
	ProductName string  `json:"productName" gorm:"column:product_name;type:varchar(100)"`
	Brand       string  `json:"brand" gorm:"column:brand;type:varchar(100)"`
	Price       float64 `json:"price" gorm:"column:price"`
}

type HistoryCart struct {
	ID         uint    `json:"idCart" gorm:"column:id_cart;primaryKey;autoIncrement;"`
	IDOrder    uint    `json:"idOrder" gorm:"column:id_order;"`
	IDCustomer uint    `json:"idCustomer" gorm:"column:id_user;"`
	IDProduct  uint    `json:"idProduct" gorm:"column:id_product;"`
	Product    Product `json:"product" gorm:"foreignKey:IDProduct"`
	Price      float64 `json:"price" gorm:"column:price"`
	Qty        int     `json:"quantity" gorm:"column:qty"`
	CreatedAt  string  `json:"created_at" gorm:"column:created_at;type:datetime"`
}

type CartResponse struct {
	ID         uint    `json:"idCart" gorm:"column:id_cart;primaryKey;autoIncrement;"`
	IDOrder    uint    `json:"-" gorm:"column:id_order;"`
	IDCustomer uint    `json:"-" gorm:"column:id_user;"`
	IDProduct  uint    `json:"-" gorm:"column:id_product;"`
	Product    Product `json:"product" gorm:"foreignKey:IDProduct"`
	Price      float64 `json:"price" gorm:"column:price"`
	Qty        int     `json:"quantity" gorm:"column:qty"`
	CreatedAt  string  `json:"-" gorm:"column:created_at;type:datetime"`
}

type AllOrder struct {
	ID         uint `json:"id" gorm:"column:id_order;primaryKey;autoIncrement;"`
	CustomerID uint `json:"customerID" gorm:"column:id_user"`
	Customer   struct {
		ID        uint   `json:"idUser" gorm:"column:id_user"`
		FullName  string `json:"fullName" gorm:"column:full_name;"`
		Gender    string `json:"gender" gorm:"column:gender;"`
		Birthdate string `json:"birthDate" gorm:"column:birth_date"`
		Username  string `json:"userName" gorm:"column:user_name;"`
	} `json:"dataCustomer"`
	Cart        []CartResponse `json:"cart" gorm:"foreignKey:IDOrder"`
	CreatedAt   string         `json:"created_at" gorm:"column:created_at;type:datetime"`
	UpdatedAt   string         `json:"updated_at" gorm:"column:updated_at;type:datetime"`
	TotalAmount float64        `json:"totalAmount" gorm:"column:total_amount"`
}

func (Order) TableName() string {
	return "order"
}

func (Product) TableName() string {
	return "product"
}

func (HistoryCart) TableName() string {
	return "history_cart"
}

func (CartResponse) TableName() string {
	return "history_cart"
}

type RequestOrder struct {
	Search string `json:"search"`
	Length int    `json:"length"`
	Start  int    `json:"start"`
}

func (o *Order) InsertOrder(db *gorm.DB) error {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	o.CreatedAt = currentTime
	o.UpdatedAt = currentTime

	return db.Create(o).Error
}

func (o *Order) UpdateOrder(db *gorm.DB) error {
	return db.Model(Order{}).Omit("id_order").Where("id_order = ?", o.ID).Updates(o).Error
}

func CountAllOrder(db *gorm.DB) int64 {
	var count int64
	result := db.Table("order").Count(&count)

	if result.Error != nil {
		return 0
	} else {
		return count
	}
}

func GetAllOrder(db *gorm.DB, limit, skip int) (error, []AllOrder) {
	var orders []AllOrder

	// Preload "Cart" association including "Product" for each "AllOrder"
	result := db.Table("order").Preload("Cart", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Product")
	}).Limit(limit).Offset(skip).Find(&orders)

	for i, order := range orders {
		user := order.Customer
		result = db.Raw("SELECT id_user, full_name, gender, birth_date, user_name from `user` where id_user = ?", order.CustomerID).First(&user)

		orders[i].Customer.ID = user.ID
		orders[i].Customer.FullName = user.FullName
		orders[i].Customer.Username = user.Username
		orders[i].Customer.Gender = user.Gender
		orders[i].Customer.Birthdate = user.Birthdate
	}

	if result.Error != nil {
		return result.Error, nil
	}

	return nil, orders
}

func GetOrderById(db *gorm.DB, id int) (error, *AllOrder) {
	var order AllOrder

	// Preload "Cart" association including "Product" for each "AllOrder"
	result := db.Table("order").Where("id_order = ?", id).Preload("Cart", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Product")
	}).First(&order)

	user := order.Customer
	result = db.Raw("SELECT id_user, full_name, gender, birth_date, user_name from `user` where id_user = ?", order.CustomerID).First(&user)

	order.Customer.ID = user.ID
	order.Customer.FullName = user.FullName
	order.Customer.Username = user.Username
	order.Customer.Gender = user.Gender
	order.Customer.Birthdate = user.Birthdate

	if result.Error != nil {
		return result.Error, nil
	}

	return nil, &order
}

func (o *Order) Delete(db *gorm.DB, id int) error {
	tx := db.Begin()

	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Where("id_order = ?", id).Delete(&HistoryCart{}).Error; err != nil {
		// Rollback the transaction if there's an error
		tx.Rollback()
		return err
	}

	// Delete from order table
	if err := tx.Where("id_order = ?", id).Delete(&Order{}).Error; err != nil {
		// Rollback the transaction if there's an error
		tx.Rollback()
		return err
	}

	// Commit the transaction if everything is successful
	return tx.Commit().Error
}
