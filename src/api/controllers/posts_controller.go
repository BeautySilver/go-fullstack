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
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err !=nil{
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo:= crud.NewRepositoryPostsCRUD(db)

	func(postRepository repository.PostRepository){
		posts, err := postRepository.FindAll()
		if err !=nil{
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		//w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		responses.JSON(w, http.StatusOK, posts)
	}(repo)

}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err !=nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post := models.Post{}
	err = json.Unmarshal(body, &post)
	if err !=nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	post.Prepare()
	err =post.Validate()
	if err !=nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	uid, err:= auth.ExtractID(r)
	if err !=nil{
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	if uid != post.AuthorID{
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	db, err := database.Connect()
	if err !=nil{
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}


	defer db.Close()

	repo:= crud.NewRepositoryPostsCRUD(db)

	func(postsRepository repository.PostRepository){
		post, err = postsRepository.Save(post)
		if err !=nil{
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, post.ID))
		responses.JSON(w, http.StatusCreated, post)
	}(repo)

}

func GetPost(w http.ResponseWriter, r *http.Request) {

	vars :=mux.Vars(r)
	pid, err :=strconv.ParseUint(vars["id"], 10, 64)
	db, err := database.Connect()
	if err !=nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	defer db.Close()

	repo:= crud.NewRepositoryPostsCRUD(db)

	func(postRepository repository.PostRepository){
		post, err := postRepository.FindById(pid)
		if err !=nil{
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		responses.JSON(w, http.StatusOK, post)
	}(repo)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars :=mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err !=nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err !=nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post := models.Post{}
	err = json.Unmarshal(body, &post)
	if err !=nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post.Prepare()
	err = post.Validate()
	if err !=nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	db, err := database.Connect()
	if err !=nil{
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo:= crud.NewRepositoryPostsCRUD(db)

	func(postRepository repository.PostRepository){
		post, err := postRepository.Update(uint64(pid), post)
		if err !=nil{
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		responses.JSON(w, http.StatusOK, post)
	}(repo)
}

func DeletPost(w http.ResponseWriter, r *http.Request) {
	vars :=mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err !=nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	uid, err:= auth.ExtractID(r)
	if err !=nil{
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	fmt.Println("User:", uid)

	db, err := database.Connect()
	if err !=nil{
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo:= crud.NewRepositoryPostsCRUD(db)

	func(postRepository repository.PostRepository){
		_, err := postRepository.Delete(pid, uid)
		if err !=nil{
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Entity", fmt.Sprintf("%d", pid))
		responses.JSON(w, http.StatusNoContent, "")
	}(repo)

}