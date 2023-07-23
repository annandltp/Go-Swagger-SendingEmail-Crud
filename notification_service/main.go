package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

type MailerRequest struct {
	BuyerEmail   string `json:"buyer_email"`
	BuyerAddress string `json:"buyer_address"`
	ProductName  string `json:"product_name"`
}

func main() {
	s := gin.Default()

	s.POST("/mailer", func(ctx *gin.Context) {
		var mailerBody MailerRequest

		if err := ctx.ShouldBindJSON(&mailerBody); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}

		m := gomail.NewMessage()
		m.SetHeader("From", "alex@example.com")
		m.SetHeader("To", mailerBody.BuyerEmail)
		m.SetHeader("Subject", "Hello!")
		m.SetBody("text/html", fmt.Sprintf("Order produk %s berhasil, dan akan dikirim ke %s segera", mailerBody.ProductName))

		d := gomail.NewDialer("smtp-relay.sendinblue.com", 587, "anan.letcol123@gmail.com", "xsmtpsib-acaf0696b5c07aaf59684b99e526d9516b2f755e55aca34b8c605b1d99415509-vNSgBZWR5qTdzFw4")

		// Send the email to Bob, Cora and Dan.
		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}

		fmt.Println(mailerBody, " <------ Mailer Body")

		ctx.JSON(http.StatusOK, "Mailer Service: Success!")
	})

	s.Run(":8001")
}
