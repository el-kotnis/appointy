package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/el-kotnis/appointy/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader((http.StatusNotFound)) //send error code to postman
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}

	if err := uc.session.DB("appointy").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(u) //converting to json
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{} //assigned with the User structure

	json.NewDecoder(r.Body).Decode(&u)             //decode the request which was in json into struct
	u.Id = bson.NewObjectId()                      //create new Id
	uc.session.DB("appointy").C("users").Insert(u) //insert into database
	uj, err := json.Marshal(u)                     //convert to json
	if err != nil {
		fmt.Println(err) //return error
	}
	//send to postman if no error
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "%s\n", uj)
}
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("appointy").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted User", oid, "\n")
}
