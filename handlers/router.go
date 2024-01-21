package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
)

type Context struct {
	Gin       *gin.Engine
	DB        *gorm.DB
	Validator *validator.Validate
}

type DefaultResponse struct {
	Error   bool        `json:"error"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (c *Context) API(group string) {
	public := c.Gin.Group(group)
	{
		public.GET("/customer", c.GetCustomers)
		public.GET("/customer/:id", c.GetCustomerById)

		public.GET("/order", c.GetOrders)
		public.GET("/order/:id", c.GetOrderById)

		public.POST("/login", c.Login)
		public.POST("/register", c.Register)
	}

	private := c.Gin.Group(group)
	{
		private.Use(func(ctx *gin.Context) {
			message, failed := CheckToken(ctx)
			resp := DefaultResponse{}

			if failed {
				resp.Message = "Please login"
				resp.Error = true
				ctx.JSON(http.StatusUnauthorized, resp)
				ctx.Abort() // Stop processing further middleware and handlers
				return
			}

			if message == "forbidden" {
				resp.Message = "You're not admin"
				resp.Error = true
				ctx.JSON(http.StatusForbidden, resp)
				ctx.Abort() // Stop processing further middleware and handlers
				return
			}
			// Continue with the next middleware/handler
			ctx.Next()
		})
		private.POST("/customer", c.InsertCustomer)
		private.PUT("/customer", c.UpdateCustomer)
		private.DELETE("/customer/:id", c.DeleteCustomer)

		private.POST("/order", c.InsertOrder)
		private.PUT("/order", c.UpdateOrder)
		private.DELETE("/order/:id", c.DeleteOrder)
	}
}
