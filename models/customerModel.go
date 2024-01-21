package models

import (
	"errors"
	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
)

// Customer struct to make structure on database
type Customer struct {
	ID        uint     `json:"idUser" gorm:"column:id_user;primaryKey;autoIncrement;"`
	FullName  string   `json:"fullName" gorm:"column:full_name;type:varchar(50)"`
	Gender    string   `json:"gender" gorm:"column:gender;type:varchar(20)"`
	Birthdate string   `json:"birthDate" gorm:"column:birth_date;type:date"`
	CreatedAt string   `json:"createdAt" gorm:"column:created_at;type:datetime"`
	UpdatedAt string   `json:"updatedAt" gorm:"column:updated_at;type:datetime"`
	UserRole  uint     `json:"userRole" gorm:"column:user_role"`
	Role      UserRole `json:"-" gorm:"foreignKey:UserRole"`
	Username  string   `json:"userName" gorm:"column:user_name;type:varchar(50)"`
	Password  string   `json:"password" gorm:"column:password;type:varchar(200)"`
}

type DataAuth struct {
	Username string `json:"userName"`
	Password string `json:"password"`
	UserRole uint   `json:"userRole"`
	IDUser   int    `json:"idUser"`
	jwt.StandardClaims
}

type DetailCustomer struct {
	ID        uint   `json:"idUser" gorm:"column:id_user;primaryKey;autoIncrement;"`
	FullName  string `json:"fullName" gorm:"column:full_name;type:varchar(50)"`
	Gender    string `json:"gender" gorm:"column:gender;type:varchar(20)"`
	Birthdate string `json:"birthDate" gorm:"column:birth_date;type:date"`
	CreatedAt string `json:"createdAt" gorm:"column:created_at;type:datetime"`
	UpdatedAt string `json:"updatedAt" gorm:"column:updated_at;type:datetime"`
	Role      Role   `json:"role" gorm:"embedded;embeddedPrefix:role_"`
	Username  string `json:"userName" gorm:"column:user_name;type:varchar(50)"`
	Password  string `json:"password" gorm:"column:password;type:varchar(200)"`
}

type Role struct {
	UserRole uint   `json:"userRole" gorm:"column:user_role"`
	Role     string `json:"role" gorm:"column:role"`
}

func (Customer) TableName() string {
	return "user"
}

type Customers []Customer

type RequestCustomer struct {
	Search string `json:"search"`
	Length int    `json:"length"`
	Start  int    `json:"start"`
}

type UserRole struct {
	ID   uint   `json:"idRole" gorm:"column:id_role;primaryKey;autoIncrement;"`
	Role string `json:"Role" gorm:"column:role;type:varchar(20);"`
}

func (UserRole) TableName() string {
	return "role"
}

func (c *Customers) GetAll(db *gorm.DB, limit, skip int, search string) error {
	if search != "" {
		return db.Model(Customer{}).
			Where("full_name LIKE ?", "%"+search+"%").
			Or("user_name LIKE ?", "%"+search+"%").
			Where("user_role = ?", 1).
			Limit(limit).Offset(skip).Find(c).Error
	}
	return db.Model(Customer{}).Where("user_role = ?", 1).Limit(limit).Offset(skip).Find(c).Error
}

func (c *DetailCustomer) GetById(db *gorm.DB, id int) error {
	return db.Model(Customer{}).
		Select("user.*, role.id_role AS role_user_role, role.role AS role_role").
		Joins("JOIN role ON user.user_role = role.id_role").
		Where("user.id_user = ?", id).
		First(c).Error
}

func (c *Customers) CountData(db *gorm.DB, req RequestCustomer) int64 {
	var count int64
	result := db.Model(Customer{})

	if req.Search != "" {
		result = result.Where("full_name LIKE ? OR user_name LIKE ?", "%"+req.Search+"%", "%"+req.Search+"%")
	}

	result = result.Where("user_role = ?", 1).Count(&count)

	if result.Error != nil {
		return 0
	}

	return count
}
func (c *Customer) Insert(db *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	c.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	c.Password = string(hashedPassword)
	//remove spaces in username
	c.Username = strings.TrimSpace(c.Username)
	return db.Model(Customer{}).Create(c).Error
}

func (c *DetailCustomer) Update(db *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	c.Password = string(hashedPassword)
	return db.Model(Customer{}).Omit("id_user").Where("id_user = ?", c.ID).Updates(c).Error
}

func (c *Customer) Delete(db *gorm.DB, id int) error {
	return db.Model(Customer{}).Where("id_user = ?", c.ID).Delete(c).Error
}

func (d *Customer) CheckUserName(db *gorm.DB, username string) bool {
	err := db.Model(Customer{}).
		Where("user_name LIKE ?", "%"+username+"%").
		First(d).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
		log.Println(err.Error())
		return true
	}
	return true
}

func (d *DataAuth) LoginCheck(db *gorm.DB, username, password string) (error, string) {
	cm := Customer{}
	err := db.Model(Customer{}).
		Where("user_name LIKE ?", "%"+username+"%").
		First(&cm).Error

	if err != nil {
		return err, ""
	}
	err = VerifyPassword(password, d.Password)
	if err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return err, ""
	}
	d.IDUser = int(cm.ID)
	d.UserRole = cm.UserRole
	err, token := CreateToken(d)
	if err != nil {
		return err, ""
	}
	return nil, token
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func CreateToken(dataAuth *DataAuth) (error, string) {
	APPLICATION_NAME := "DBO_JWT"
	LOGIN_EXPIRATION_DURATION := 7 * time.Hour
	JWT_SIGNING_METHOD := jwt.SigningMethodHS256
	JWT_SIGNATURE_KEY := []byte("dbo2024@jakarta")

	dataAuth.StandardClaims = jwt.StandardClaims{
		Issuer:    APPLICATION_NAME,
		ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
	}
	dataAuth.Password = "*****"
	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		dataAuth,
	)
	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return err, ""
	}

	return nil, signedToken
}
