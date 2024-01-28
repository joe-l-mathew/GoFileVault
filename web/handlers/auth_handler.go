package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joe-l-mathew/GoFileVault/models"
	"github.com/joe-l-mathew/GoFileVault/pkg"
	"github.com/joe-l-mathew/GoFileVault/pkg/db"
	"github.com/joe-l-mathew/GoFileVault/utils"
	"gorm.io/gorm"
)

type UserCreateAccount struct {
	Name     string `json:"name"`
	EmailId  string `json:"email"`
	Password string `json:"password"`
}

type UserSignin struct {
	EmailId  string `json:"email"`
	Password string `json:"password"`
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var signupData UserCreateAccount
	err = json.Unmarshal(body, &signupData)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	encrypetedPassword, hashErr := utils.EncryptPassword(signupData.Password)
	if hashErr != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	model := models.User{
		Name:     signupData.Name,
		Password: encrypetedPassword,
		Email:    signupData.EmailId,
	}

	if err = db.DB.Create(&model).Error; err != nil {
		http.Error(w, "Error creating user", http.StatusBadRequest)
		return
	}
	os.MkdirAll(fmt.Sprint(pkg.StorageName, "/", model.ID), 0755)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Succesfuly created user"))

}

func SignInUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	var signinUser UserSignin
	w.Header().Set("Content-Type", "application/json")

	err = json.Unmarshal(body, &signinUser)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	var user models.User
	if err := db.DB.Where("email=?", signinUser.EmailId).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "No user found", http.StatusNotFound)
		} else {
			http.Error(w, "internal Error", http.StatusBadRequest)
		}
		return
	}
	if utils.VerifyPassword(signinUser.Password, user.Password) {
		w.WriteHeader(http.StatusOK)
		token, _ := utils.GenerateToken(user.ID)
		data := map[string]interface{}{
			"data":  user,
			"token": token,
		}
		jsonData, _ := json.Marshal(data)
		w.Write(jsonData)

	} else {
		http.Error(w, "Invalid password", http.StatusBadRequest)
	}
}
