package edc

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UniqueDBRule(fl validator.FieldLevel, c *fiber.Ctx) bool {
	var count int
	params := strings.Split(fl.Param(), ":")

	if len(params) < 2 {
		panic("UniqueDBRule: invalid parameters")
	}

	tableName := params[0]
	columnName := params[1]

	query := fmt.Sprintf("SELECT COUNT(1) FROM %s WHERE %s = ?", tableName, columnName)

	if len(params) > 2 {
		exceptValue := c.Params(params[2])
		exceptColumn := "id"

		if len(params) > 3 {
			exceptColumn = params[3]
		}

		query = fmt.Sprintf("%s AND %s != '%s'", query, exceptColumn, exceptValue)
	}

	Edc.DB.Raw(query, fl.Field()).Scan(&count)

	return count == 0
}

func ExistsDBRule(fl validator.FieldLevel, c *fiber.Ctx) bool {
	var count int
	params := strings.Split(fl.Param(), ":")

	if len(params) < 2 {
		panic("ExistsDBRule: invalid parameters")
	}

	tableName := params[0]
	columnName := params[1]

	query := fmt.Sprintf("SELECT COUNT(1) FROM %s WHERE %s = ?", tableName, columnName)

	Edc.DB.Raw(query, fl.Field()).Scan(&count)

	return count > 0
}
