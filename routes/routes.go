package routes

import (
	"cobalagi/models"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	route := echo.New()

	//API
	route.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Halo Kesayangan, wanita impian, tujuan kahir, rumah untukku pulang dan pujaan hati. Love you so much <3 jangan lupa makan")
	})

	route.POST("user/create_user", func(c echo.Context) error {
		user := new(models.Users)
		c.Bind(user)
		contentType := c.Request().Header.Get("Content-Type")
		if contentType == "application/json" {
			fmt.Println("Request dari json")
		} else if strings.Contains(contentType, "multipart/form-data") || contentType == "application/x-www-form-urlencoded" {
			file, err := c.FormFile("ktp")
			if err != nil {
				fmt.Println("ktp kosong")
			} else {
				src, err := file.Open()
				if err != nil {
					return err
				}
				defer src.Close()
				dst, err := os.Create(file.Filename)
				if err != nil {
					return err
				}
				defer dst.Close()
				if _, err = io.Copy(dst, src); err != nil {
					return err
				}

				user.Ktp = file.Filename
				fmt.Println("fileada, akan disimpan")
			}
		}
		response := new(models.Response)
		if user.CreateUser() != nil {
			response.ErrorCode = 10 //method create
			response.Message = "Gagal menambahkan data user"
		} else {
			response.ErrorCode = 0
			response.Message = "Sukses menambahkan data user"
			response.Data = *user
		}
		return c.JSON(http.StatusOK, response)
	})

	route.PUT("user/update_user/:email", func(c echo.Context) error {
		user := new(models.Users)
		c.Bind(user)
		response := new(models.Response)
		if user.UpdateUser(c.Param("email")) != nil { //method update
			response.ErrorCode = 10
			response.Message = "Gagal update user"

		} else {
			response.ErrorCode = 0
			response.Message = "Sukses update data user"
			response.Data = *user
		}

		return c.JSON(http.StatusOK, response)
	})

	route.DELETE("user/delete_user/:email", func(c echo.Context) error {
		user, _ := models.GetOneByEmail(c.Param("email")) //method getby email
		response := new(models.Response)
		if user.DeleteUser() != nil { //method hapus user
			response.ErrorCode = 10
			response.Message = "Gagal meghapus data"
		} else {
			response.ErrorCode = 0
			response.Message = "Berhasil menghapus data"
		}

		return c.JSON(http.StatusOK, response)
	})

	route.GET("user/search_user", func(c echo.Context) error {
		response := new(models.Response)
		users, err := models.GetAll(c.QueryParam("keywords")) //method get all
		if err != nil {
			response.ErrorCode = 10
			response.Message = "Gagal melihat data user"
		} else {
			response.ErrorCode = 0
			response.Message = "Berhasil melhat data"
			response.Data = users
		}
		return c.JSON(http.StatusOK, response)
	})
	return route
}
