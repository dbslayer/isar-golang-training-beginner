package api

import (
	"fmt"
	"isar-golang-training-beginner/db"
	"isar-golang-training-beginner/model"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

const (
	ActiveStatus   = "ACTIVE"
	IncativeStatus = "INACTIVE"
	ExpiredStatus  = "EXPIRED"
)

func GetPaymentCodes(c echo.Context) error {
	payment_codes := c.Param("payment_codes")
	db := db.DbManager()
	payments := model.PaymentCodes{}
	db.First(&payments, "id=?", payment_codes)

	if payments.ID == uuid.Nil {
		return c.JSON(http.StatusNotFound, "What are you looking for?")
	}
	//spew.Dump(json.Marshal(payments))
	return c.JSON(http.StatusOK, payments)
}

func PostPaymentCodes(c echo.Context) error {
	p := new(model.PaymentCodes)
	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db := db.DbManager()

	currentTime := time.Now()
	p.Status = ActiveStatus
	p.CreatedDate = currentTime

	expirationDate := time.Now().AddDate(50, 0, 1) // create expiration date to next 50 years + 1 day
	p.ExpirationDate = expirationDate.Format("2006-01-02")

	status := db.Create(&p)
	fmt.Println(status)
	return c.JSON(http.StatusCreated, p)
}
