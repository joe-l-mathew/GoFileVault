package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/joe-l-mathew/GoFileVault/models"
	"github.com/joe-l-mathew/GoFileVault/pkg/db"
)

func GetUserFiles(w http.ResponseWriter, r *http.Request) {
	var userFiles []models.UserFiles
	if err := db.DB.Find(&userFiles).Error; err != nil {
		http.Error(w, "Error getting data", http.StatusInternalServerError)
		return
	}
	returnData := map[string]interface{}{
		"files": userFiles,
	}
	data, err := json.Marshal(returnData)
	if err != nil {
		http.Error(w, "Error marshaling data", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
