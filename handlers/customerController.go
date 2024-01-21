package handlers

import (
	"dbo-test/helpers/env"
	"dbo-test/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (c *Context) GetCustomers(ctx *gin.Context) {
	request := models.RequestCustomer{}
	request.Start, _ = strconv.Atoi(ctx.Query("start"))
	request.Length, _ = strconv.Atoi(ctx.Query("length"))
	request.Search = ctx.Query("search")

	resp := DefaultResponse{}

	data := models.Customers{}
	if request.Start <= 0 {
		request.Start = 1
	}
	if request.Length <= 0 {
		request.Length = env.Get().DefaultLimit
	}

	skip := (request.Start - 1) * request.Length

	err := data.GetAll(c.DB, request.Length, skip, request.Search)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err.Error())
			resp.Error = false
			resp.Message = "Data not found"
			resp.Data = nil
			ctx.JSON(http.StatusNotFound, resp)
			return
		} else {
			log.Println(err.Error())
			resp.Error = true
			resp.Message = err.Error()
			resp.Data = nil
			ctx.JSON(http.StatusInternalServerError, resp)
			return
		}
	}

	if len(data) == 0 {
		resp.Data = nil
		resp.Error = false
		resp.Message = "Data not found"
		ctx.JSON(http.StatusNotFound, resp)
		return
	}

	Alldata := data.CountData(c.DB, request)

	for i := range data {
		data[i].Password = "*****"
	}

	metadata := gin.H{
		"currentPage": request.Start,
		"totalPage":   int(math.Ceil(float64(Alldata) / float64(request.Length))),
		"dataOnPage":  len(data),
		"totalData":   Alldata,
	}

	response := gin.H{
		"data":     data,
		"error":    false,
		"message":  "Success",
		"metadata": metadata,
	}

	ctx.JSON(http.StatusOK, response)
	return

}

func (c *Context) GetCustomerById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	resp := DefaultResponse{}
	if id == 0 {
		resp.Error = false
		resp.Message = "invalid input parameter id"
		resp.Data = nil

		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	var data models.DetailCustomer
	err := data.GetById(c.DB, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err.Error())
			resp.Error = false
			resp.Message = "Data not found"
			resp.Data = nil
			ctx.JSON(http.StatusNotFound, resp)
			return
		} else {
			log.Println(err.Error())
			resp.Error = true
			resp.Message = err.Error()
			resp.Data = nil
			ctx.JSON(http.StatusInternalServerError, resp)
			return
		}
	}
	data.Password = "*****"

	resp.Data = data
	resp.Error = false
	resp.Message = "Success"

	ctx.JSON(http.StatusOK, resp)
	return
}

func (c *Context) InsertCustomer(ctx *gin.Context) {
	resp := DefaultResponse{}
	data := models.Customer{}
	err := ctx.Bind(&data)

	if err != nil {
		log.Println(err.Error())
		resp.Error = true
		resp.Message = err.Error()
		resp.Data = nil
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	err = data.Insert(c.DB)

	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		if mysqlErr.Number == 1452 { // MySQL error code for foreign key constraint violation
			log.Println(err.Error())
			resp.Error = true
			resp.Message = "Please add role user at table 'role' firstly"
			ctx.JSON(http.StatusInternalServerError, resp)
			return
		}
	}

	if err != nil {
		log.Println(err.Error())
		resp.Error = true
		resp.Message = err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Data = nil
	resp.Error = false
	resp.Message = "Success"

	ctx.JSON(http.StatusCreated, resp)
	return

}

func (c *Context) UpdateCustomer(ctx *gin.Context) {
	rqdata := models.Customer{}
	resp := DefaultResponse{}
	if err := ctx.Bind(&rqdata); err != nil {
		log.Println(err.Error())
		resp.Error = true
		resp.Message = err.Error()
		resp.Data = nil
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	// cek data
	var dbdata models.DetailCustomer
	err := dbdata.GetById(c.DB, int(rqdata.ID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err.Error())
			resp.Error = false
			resp.Message = "Data not found"
			resp.Data = nil
			ctx.JSON(http.StatusNotFound, resp)
			return
		} else {
			log.Println(err.Error())
			resp.Error = true
			resp.Message = err.Error()
			resp.Data = nil
			ctx.JSON(http.StatusInternalServerError, resp)
			return
		}
	}

	dbdata.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	dbdata.Password = rqdata.Password
	dbdata.Username = rqdata.Username
	dbdata.Birthdate = rqdata.Birthdate
	dbdata.FullName = rqdata.FullName
	dbdata.Gender = rqdata.Gender

	err = dbdata.Update(c.DB)
	if err != nil {
		log.Println(err.Error())
		resp.Error = true
		resp.Message = err.Error()
		resp.Data = nil
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Data = nil
	resp.Error = false
	resp.Message = "Success"

	ctx.JSON(http.StatusOK, resp)
	return

}

func (c *Context) DeleteCustomer(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	resp := DefaultResponse{}
	if id == 0 {
		resp.Error = false
		resp.Message = "invalid input parameter id"
		resp.Data = nil

		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	data := models.DetailCustomer{}
	err := data.GetById(c.DB, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err.Error())
			resp.Error = false
			resp.Message = "Data not found"
			resp.Data = nil
			ctx.JSON(http.StatusNotFound, resp)
			return
		} else {
			log.Println(err.Error())
			resp.Error = true
			resp.Message = err.Error()
			resp.Data = nil
			ctx.JSON(http.StatusInternalServerError, resp)
			return
		}
	}

	var cm models.Customer
	err = cm.Delete(c.DB, int(data.ID))
	if err != nil {
		log.Println(err.Error())
		resp.Error = true
		resp.Message = err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Data = nil
	resp.Error = false
	resp.Message = "Success"

	ctx.JSON(http.StatusOK, resp)
	return
}

func (c *Context) Register(ctx *gin.Context) {
	data := models.DataAuth{}
	err := ctx.Bind(&data)
	resp := DefaultResponse{}
	if err != nil {
		log.Println(err.Error())
		resp.Error = true
		resp.Message = err.Error()
		resp.Data = nil
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	cm := models.Customer{}
	isUsed := cm.CheckUserName(c.DB, data.Username)
	if isUsed {
		resp.Error = true
		resp.Message = "username telah digunakan"
		resp.Data = nil
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	user := models.Customer{}
	user.Username = data.Username
	user.Password = data.Password
	user.UserRole = data.UserRole
	err = user.Insert(c.DB)

	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		if mysqlErr.Number == 1452 { // MySQL error code for foreign key constraint violation
			log.Println(err.Error())
			resp.Error = true
			resp.Message = "Please add role user at table 'role' firstly"
			ctx.JSON(http.StatusInternalServerError, resp)
			return
		}
	}

	if err != nil {
		log.Println(err.Error())
		resp.Error = true
		resp.Message = err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Data = nil
	resp.Error = false
	resp.Message = "Success Register"

	ctx.JSON(http.StatusCreated, resp)
	return

}

func (c *Context) Login(ctx *gin.Context) {
	data := models.DataAuth{}
	err := ctx.Bind(&data)
	resp := DefaultResponse{}
	if err != nil {
		log.Println(err.Error())
		resp.Error = true
		resp.Message = err.Error()
		resp.Data = nil
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	//get id by username
	err, token := data.LoginCheck(c.DB, data.Username, data.Password)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err.Error())
			resp.Error = false
			resp.Message = "Username or Password is wrong"
			resp.Data = nil
			ctx.JSON(http.StatusNotFound, resp)
			return
		} else {
			log.Println(err.Error())
			resp.Error = true
			resp.Message = err.Error()
			resp.Data = nil
			ctx.JSON(http.StatusInternalServerError, resp)
			return
		}
	}

	data.Password = "*****"
	resp.Error = true
	resp.Message = token
	resp.Data = data
	ctx.JSON(http.StatusOK, resp)
	return

}

func CheckToken(ctx *gin.Context) (message string, status bool) {
	JWT_SIGNING_METHOD := jwt.SigningMethodHS256
	JWT_SIGNATURE_KEY := []byte("dbo2024@jakarta")
	authorizationHeader := ctx.Request.Header.Get("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		return "Invalid token", true
	}
	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("signing method invalid")
		}
		return JWT_SIGNATURE_KEY, nil
	})
	if err != nil {
		return "Error", true
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "Error", true
	}

	// Access the "userInfo" claim from the claims
	userInfoClaim := claims["userRole"].(float64)

	if userInfoClaim != 2 && userInfoClaim != 1 {
		return "forbidden", true
	}

	return "success", false
}
