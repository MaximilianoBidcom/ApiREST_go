package mysql

import (
	"fmt"
	"log"

	"../../models"
)

type GlobalImplMysql struct{}

func checkErr(model string, accion string, err error) {
	if err != nil {
		log.Println("No se pudo "+accion+" "+model+". Error desde IMPL: ", err)
	}
}

//TODO: Agregar los DB().BEGIN() y COMMIT
func (dao GlobalImplMysql) Create(x *models.GlobalModel, model string) (models.GlobalModel, error) {
	switch model {
	case "post":
		err := DB().CreateAndRead(&x.Post)
		checkErr("post", "crear", err)
	default:
		log.Println("Modelo ingresado no existente")
	}
	return *x, nil
}
func (dao GlobalImplMysql) GetAll(model string) (models.GlobalModels, error) {
	var a models.GlobalModels

	switch model {
	case "post":
		err := DB().Read(&a.Actividad, "SELECT * FROM post")
		checkErr("post", "getAll", err)
	default:
		log.Println("Modelo ingresado no existente")
	}
	return a, nil
}
func (dao GlobalImplMysql) GetByID(id int, model string) (models.GlobalModel, error) {
	var x models.GlobalModel

	switch model {
	case "post":
		err := DB().Read(&x.Continente, "SELECT * FROM post WHERE id = ?", id)
		checkErr("post", "getByID", err)
	default:
		log.Println("Modelo ingresado no existente")
	}
	return x, nil
}
func (dao GlobalImplMysql) GetBy(x models.GlobalModel, model string) (models.GlobalModels, error) {
	a := models.GlobalModels{}
	switch model {
	case "post":
		err := DB().Read(&a.post, "SELECT * FROM post WHERE post = ?", x.Post.Post)
		checkErr("post", "getOne", err)
	default:
		log.Println("Modelo ingresado no existente")
	}
	fmt.Println("MSJ: ", a)
	return a, nil
}
func (dao GlobalImplMysql) Update(x models.GlobalModel, model string) (models.GlobalModel, error) {
	switch model {
	case "post":
		var pos models.Post

		err := DB().Read(&pos, "SELECT * FROM post WHERE id = ?", x.Post.ID)
		checkErr("post", "getByID", err)

		if x.Post.Post == "" {
			x.Post.Post = pos.Imagen
		}
		err = DB().Update(x.Post)
		checkErr("post", "actualizar", err)
	default:
		log.Println("Modelo ingresado no existente")
	}
	return x, nil
}
func (dao GlobalImplMysql) Delete(x *models.GlobalModel, model string) (string, error) {
	var err error
	var msjReturn = ""
	switch model {
	case "post":
		msjReturn = "Post eliminado correctamente."
		err = DB().Delete(x.Post)
		checkErr("imagen", "eliminar", err)
	default:
		msjReturn = "Modelo ingresado no 	existente"
	}
	return msjReturn, err
}
