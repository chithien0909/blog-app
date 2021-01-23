package controllers

import (
	"../database"
	"../models"
	"../repository"
	"../repository/crud"
	"../responses"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetPosts(w http.ResponseWriter, _ *http.Request) {
	db, err := database.Connect()

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryPostsCRUD(db)

	func (postsRepository repository.PostRepository) {
		posts, err := postsRepository.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusOK, posts)
	}(repo)

}
func CreatePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post := models.Post{}

	err = json.Unmarshal(body, &post)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	post.Prepare()
	err = post.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	db, err := database.Connect()

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryPostsCRUD(db)

	func (postsRepository repository.PostRepository) {
		post, err = postsRepository.Save(post)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, post.ID))
		responses.JSON(w, http.StatusCreated, post)
	}(repo)
}
func GetPost(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryPostsCRUD(db)

	func (postsRepository repository.PostRepository) {
		post, err := postsRepository.FindById(uid)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusOK, post)
	}(repo)
}
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post := models.Post{}
	err = json.Unmarshal(body, &post)
	post.Prepare()
	err = post.Validate("update")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryPostsCRUD(db)
	func (postsRepository repository.PostRepository) {
		rowEffect, err := postsRepository.Update(uid, post)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusOK, rowEffect)
	}(repo)
}
func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryPostsCRUD(db)
	func (postsRepository repository.PostRepository) {
		_, err := postsRepository.Delete(uid)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Entity", fmt.Sprintf("%d", uid))
		responses.JSON(w, http.StatusNoContent, "")
	}(repo)
}
