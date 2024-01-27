package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/joe-l-mathew/GoFileVault/models"
	"github.com/joe-l-mathew/GoFileVault/pkg/db"
)

type UserCreateAccount struct {
	Name     string `json:"name"`
	EmailId  string `json:"email"`
	Password string `json:"password"`
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println("call reached here")

	model := models.User{
		Name:     signupData.Name,
		Password: signupData.Password,
		Email:    signupData.EmailId,
	}
	if err = db.DB.Create(&model).Error; err != nil {
		http.Error(w, "Error creating user", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Succesfuly created user"))

}
