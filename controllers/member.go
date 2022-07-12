package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nogueirahy/api-gymfull/database"
	"github.com/nogueirahy/api-gymfull/models"
)

func ShowAllMembers(c echo.Context) error {
	db := database.GetDatabase()
	q := c.QueryParams()

	page, _ := strconv.Atoi(q.Get("page"))
	if page == 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(q.Get("page_size"))
	switch {
	case pageSize > 100:
		pageSize = 100
	}

	offset := (page - 1) * pageSize

	m := new([]models.Member)

	err := db.Offset(offset).Limit(pageSize).Find(&m).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, "error "+err.Error())
	}

	return c.JSON(http.StatusOK, m)
}

func FindMember(c echo.Context) error {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, "error "+err.Error())
	}

	db := database.GetDatabase()

	var member models.Member

	err = db.First(&member, newid).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, "error "+err.Error())
	}

	return c.JSON(http.StatusOK, member)
}

func CreateMember(c echo.Context) error {
	db := database.GetDatabase()
	m := new(models.Member)

	err := c.Bind(m)

	if err != nil {
		return nil
	}

	err = db.Create(&m).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, "error >> "+err.Error())
	}

	return c.JSON(http.StatusCreated, m)
}

func DeleteMember(c echo.Context) error {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, "error "+err.Error())
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Member{}, newid).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, "error "+err.Error())
	}

	return c.NoContent(http.StatusNoContent)

}

func EditMember(c echo.Context) error {
	db := database.GetDatabase()

	m := new(models.Member)

	err := c.Bind(m)

	if err != nil {
		return nil
	}

	err = db.Save(&m).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, "error "+err.Error())
	}

	return c.JSON(http.StatusOK, m)
}
