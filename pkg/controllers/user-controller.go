package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/sky-vibhu/go-userstore/pkg/utils"
	models "github.com/sky-vibhu/go-userstore/pkg/models/user"
)

var NewUser models.User

func GetUser(w http.ResponseWriter, r *http.Request){
	newUsers:= models.GetAllUsers()
	res, _ := json.Marshal(newUsers)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}



func CreateUser(w http.ResponseWriter, r *http.Request){
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	b := CreateUser.CreateUser()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userId := vars["UserId"]
	ID, err := strconv.ParseInt(userId, 0 , 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	user := models.DeleteUser(ID)  
	res, _ := json.Marshal(user)
	w.Header().Set("Content-type", "pkglication.json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.Response, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	Id, err := strconv.ParseInt(userId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userDetails, _:= models.GetUserbyId(Id)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)   
	vars := mux.Vars(r)
	UserId := vars["userId"]
	ID, err := strconv.ParseInt(UserId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userDetails, db := models.GetUserById(ID)
	if updateUser.Name != ""{	w.WriteHeader(http.StatusOK)

		userDetails.Name = updateUser.Name
	}
	if updateUser.Email != ""{
		userDetails.Email = updateUser.Email
	}
	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}




