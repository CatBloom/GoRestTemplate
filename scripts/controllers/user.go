package controllers

import (
	"log"
	"main/models"
	"main/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	List(echo.Context) error
	Get(echo.Context) error
	Post(echo.Context) error
}

type userController struct {
	m models.UserModel
}

func NewUserController(m models.UserModel) UserController {
	return &userController{m}
}

func (uc *userController) List(c echo.Context) error {
	req := types.ReqUser{}
	req.Limit = c.QueryParam("limit")
	req.Order = c.QueryParam("order")

	res, err := uc.m.GetUsers(req)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}
	return c.JSON(http.StatusOK, res)
}

func (uc *userController) Get(c echo.Context) error {
	id := c.Param("id")

	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return err
	}

	res, err := uc.m.GetUserByID(intID)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}
	return c.JSON(http.StatusOK, res)
}

func (uc *userController) Post(c echo.Context) error {
	req := types.ReqCreateUser{}

	if err := c.Bind(&req); err != nil {
		log.Printf("error:%s", err.Error())
		return err
	}

	res, err := uc.m.CreateUser(req)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}
	return c.JSON(http.StatusOK, res)
}
