package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
	"github.com/ismailozdel/micro/gateway/api"
	"github.com/ismailozdel/micro/gateway/global"
	"github.com/kelseyhightower/envconfig"
)

func main() {

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {

			code := fiber.StatusInternalServerError
			message := "Internal Server Error"

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
				message = e.Message
			} else {
				message = err.Error()
			}

			return ctx.Status(code).JSON(&global.Response{Code: 1, Message: message})
		},
	})

	//panic durumunda logla ve 500 döndür
	app.Use(func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				msg := ""
				if utf8.Valid(debug.Stack()) {
					msg += fmt.Sprintf("Stack Trace: %v", string(debug.Stack()))
				}

				log.Printf("%v \r\n %v", r, msg)
				fmt.Printf("%v \r\n %v", r, msg)
				c.Status(fiber.StatusInternalServerError).JSON(&global.Response{Code: 1, Message: fmt.Sprintf("%v", r)})
			}
		}()
		return c.Next()
	})

	invoicehost := os.Getenv("INVOICE_HOST")
	stockhost := os.Getenv("STOCK_HOST")
	userhost := os.Getenv("USER_HOST")

	invoiceGrpcConn := NewGrpcClient(invoicehost)
	defer invoiceGrpcConn.Close()
	stockGrpcConn := NewGrpcClient(stockhost)
	defer stockGrpcConn.Close()

	userGrpcConn := NewGrpcClient(userhost)
	defer userGrpcConn.Close()

	config := Config{}
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err)
	}

	api.SetupRouter(app, invoiceGrpcConn, stockGrpcConn, userGrpcConn)

	err = app.Listen(":" + config.Port)
	global.CheckError(err)
}
