package controllers

import (
	"../auth"
	"../database"
	"../models"
	"../repository"
	"../repository/crud"
	"../responses"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	db, err := database.Connect()
	if err !=nil{
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo:= crud.NewRepositoryUsersCRUD(db)

	func(usersRepository repository.UserRepository){
		users, err := usersRepository.FindAll()
		if err !=nil{
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}

		responses.JSON(w, http.StatusOK, users)
	}(repo)
	tmpl, _ := template.ParseFiles("test-client/index.html")
	tmpl.Execute(w, models.User{})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err !=nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err !=nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err =user.Validate("")
	if err !=nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err !=nil{
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo:= crud.NewRepositoryUsersCRUD(db)

	func(usersRepository repository.UserRepository){
      user, err = usersRepository.Save(user)
		if err !=nil{
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user.ID))
      responses.JSON(w, http.StatusCreated, user)
	}(repo)

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    uid, err :=strconv.ParseUint(vars["id"], 10, 32)
	if err !=nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
    user:= models.User{}
	user.Prepare()
	err =user.Validate("update")
	if err !=nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err !=nil{
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo:= crud.NewRepositoryUsersCRUD(db)

	func(usersRepository repository.UserRepository){
		user, err := usersRepository.FindById(uint32(uid))
		if err !=nil{
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		responses.JSON(w, http.StatusOK, user)
	}(repo)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err :=strconv.ParseUint(vars["id"], 10, 32)
	if err !=nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err !=nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err !=nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	tokenid, err:= auth.ExtractID(r)
	if err !=nil{
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	if tokenid != uint32(uid){
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	db, err := database.Connect()
	if err !=nil{
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo:= crud.NewRepositoryUsersCRUD(db)

	func(usersRepository repository.UserRepository){
		rows, err := usersRepository.Update(uint32(uid), user)
		if err !=nil{
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		responses.JSON(w, http.StatusOK, rows)
	}(repo)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err :=strconv.ParseUint(vars["id"], 10, 32)
	if err !=nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	tokenid, err:= auth.ExtractID(r)
	if err !=nil{
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	if tokenid != uint32(uid){
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	db, err := database.Connect()
	if err !=nil{
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo:= crud.NewRepositoryUsersCRUD(db)

	func(usersRepository repository.UserRepository){
		_, err := usersRepository.Delete(uint32(uid))
		if err !=nil{
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
        w.Header().Set("Entity", fmt.Sprintf("%d", uid))
		responses.JSON(w, http.StatusNoContent, "")
	}(repo)

}
