package edc

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validator *validator.Validate

func ValidatorMiddleware(schema interface{}, c *fiber.Ctx) error {
	Validator = validator.New(validator.WithRequiredStructEnabled())

	Validator.RegisterValidation("uniqueDB", func(fl validator.FieldLevel) bool {
		return UniqueDBRule(fl, c)
	})
	Validator.RegisterValidation("existsDB", func(fl validator.FieldLevel) bool {
		return ExistsDBRule(fl, c)
	})

	dataType := reflect.ValueOf(schema).Elem().Type()

	if err := Validator.Struct(schema); err != nil {
		var errSchema []fiber.Map

		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.StructField()
			if field, exists := dataType.FieldByName(err.StructField()); exists {
				fieldName = field.Tag.Get("json")
			}

			errSchema = append(errSchema, fiber.Map{
				"field":       fieldName,
				"rule":        err.Tag(),
				"type":        err.Kind().String(),
				"value":       err.Value(),
				"valid_value": err.Param(),
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fiber.ErrBadRequest.Message,
			"errors":  errSchema,
		})
	}

	return c.Next()
}

func GetEntityFromIDParam(
	model interface{},
	schema interface{},
	localName string,
	c *fiber.Ctx,
) error {
	id := c.Params("id")
	if res := Edc.DB.
		Model(model).
		Where("id = ?", id).
		First(&schema); res.Error != nil {
		return ErrorResponse(res.Error, c)
	}
	c.Locals(localName, schema)
	return c.Next()
}
