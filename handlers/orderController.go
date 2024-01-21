package handlers

import (
	"dbo-test/helpers/env"
	"dbo-test/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

func (c *Context) GetOrders(ctx *gin.Context) {
	request := models.RequestOrder{}
	request.Start, _ = strconv.Atoi(ctx.Query("start"))
	request.Length, _ = strconv.Atoi(ctx.Query("length"))
	request.Search = ctx.Query("search")

	resp := DefaultResponse{}

	if request.Start <= 0 {
		request.Start = 1
	}
	if request.Length <= 0 {
		request.Length = env.Get().DefaultLimit
	}

	skip := (request.Start - 1) * request.Length

	err, data := models.GetAllOrder(c.DB, request.Length, skip)
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

	Alldata := models.CountAllOrder(c.DB)

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

func (c *Context) GetOrderById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	resp := DefaultResponse{}
	if id == 0 {
		resp.Error = false
		resp.Message = "invalid input parameter id"
		resp.Data = nil

		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	err, data := models.GetOrderById(c.DB, id)
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

	resp.Data = data
	resp.Error = false
	resp.Message = "Success"

	ctx.JSON(http.StatusOK, resp)
	return
}

func (c *Context) InsertOrder(ctx *gin.Context) {
	data := models.Order{}
	resp := DefaultResponse{}
	if err := ctx.Bind(&data); err != nil {
		log.Println(err.Error())
		resp.Error = true
		resp.Message = err.Error()
		resp.Data = nil
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	data.TotalAmount = 0

	for i := range data.Cart {
		data.Cart[i].IDCustomer = data.CustomerID
		data.Cart[i].CreatedAt = time.Now().Format("2006-01-02 15:04:05")
		data.TotalAmount += data.Cart[i].Price * float64(data.Cart[i].Qty)
	}

	err := data.InsertOrder(c.DB)
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

func (c *Context) UpdateOrder(ctx *gin.Context) {
	data := models.Order{}
	resp := DefaultResponse{}
	if err := ctx.Bind(&data); err != nil {
		log.Println(err.Error())
		resp.Error = true
		resp.Message = err.Error()
		resp.Data = nil
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	//cek data terlebih dahulu
	err, order := models.GetOrderById(c.DB, int(data.ID))
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

	if data.TotalAmount == 0 {
		data.TotalAmount = order.TotalAmount
	}

	if data.CreatedAt == "" {
		data.CreatedAt = order.CreatedAt
	}

	err = data.UpdateOrder(c.DB)
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

func (c *Context) DeleteOrder(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	resp := DefaultResponse{}
	if id == 0 {
		resp.Error = false
		resp.Message = "invalid input parameter id"
		resp.Data = nil

		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	err, data := models.GetOrderById(c.DB, id)
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

	order := models.Order{}
	err = order.Delete(c.DB, int(data.ID))
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
