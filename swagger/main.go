package main

import (
	"net/http"

	_ "kbtg-boocamp-swagger/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const COUNTER_CACHE_KEY = "counter"

// Pet model info
// @Description Pet information
type Pet struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// HTTPError model info
// @Description HTTPError information
type HTTPError struct {
	BusinessErrorCode int    `json:"businessErrorCode"`
	Description       string `json:"description"`
}

var pets []Pet // ประกาศตัวแปร global เพื่อเก็บข้อมูล Pet

// @title Echo Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /
// @schemes http
func main() {

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	pets = append(pets, Pet{ID: 1, Name: "ChaNom"}, Pet{ID: 2, Name: "Mali"})
	e.GET("/pets", allPet)
	e.POST("/pets", createPet)
	e.Logger.Fatal(e.Start(":8080"))
}

// ListPet godoc
// @Summary      List of all pet
// @Description  list pets
// @Tags         pet
// @Accept       json
// @Produce      json
// @Success      200  {array}   Pet
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Router       /pets [get]
func allPet(c echo.Context) error {

	return c.JSON(http.StatusOK, pets)
}

// CreatePet godoc
// @Summary      Create a new pet
// @Description  Create a new pet with the given information
// @Tags         pet
// @Accept       json
// @Produce      json
// @Param        pet body Pet true "Pet object that needs to be added to the store"
// @Success      201  {object} Pet
// @Failure      400  {object} HTTPError
// @Failure      500  {object} HTTPError
// @Router       /pets [post]
func createPet(c echo.Context) error {
	newPet, err := bindPetFromBody(c)
	if err != nil {
		return err
	}
	pets = append(pets, *newPet)

	// ทำการบันทึก newPet ลงในฐานข้อมูลหรือทำการประมวลผลต่อไป

	// ส่ง response กลับไป
	return c.JSON(http.StatusCreated, newPet)
}

// bindPetFromBody ฟังก์ชันที่ใช้เพื่อแปลงข้อมูล JSON จาก request body เป็นโครงสร้างข้อมูล Pet
func bindPetFromBody(c echo.Context) (*Pet, error) {
	newPet := new(Pet)
	if err := c.Bind(newPet); err != nil {
		return nil, err
	}

	return newPet, nil
}
