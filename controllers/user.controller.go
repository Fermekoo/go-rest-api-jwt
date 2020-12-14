package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Fermekoo/blog-dandi/helpers"
	"github.com/Fermekoo/blog-dandi/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func FetchAllUser(c echo.Context) error {
	result, err := models.FetchAllUser()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreUser(c echo.Context) error {
	email := c.FormValue("email")
	fullname := c.FormValue("fullname")
	password := c.FormValue("password")

	hashPassword, err := helpers.HashPassword(password)

	result, err := models.StoreUser(email, fullname, hashPassword)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateUser(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("email")
	fullname := c.FormValue("fullname")

	convID, err := strconv.Atoi(id) //convert id string to Int
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateUser(convID, name, fullname)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteUser(c echo.Context) error {
	id := c.FormValue("id")

	convID, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	result, err := models.DeleteUser(convID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func CheckLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	res, err := models.CheckLogin(email, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	//generate token

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret")) //secret must be change on production

	if err != nil {
		return c.JSON(http.StatusOK, map[string]string{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"token":  t,
	})
}
