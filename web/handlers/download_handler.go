package handlers

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joe-l-mathew/GoFileVault/models"
	"github.com/joe-l-mathew/GoFileVault/pkg/db"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := mux.Vars(r)
	// Access individual parameters by key
	fileId := queryParams["filefileIdId"]
	var userFile models.UserFiles
	if err := db.DB.First(&userFile, fileId).Error; err != nil {
		http.Error(w, "DB error", http.StatusNotFound)
		return
	}
	filePath := userFile.FilePath // Replace with the actual file path

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Get the file's information
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Error getting file information", http.StatusInternalServerError)
		return
	}

	// Set the Content-Disposition header to prompt the user to download the file
	w.Header().Set("Content-Disposition", "attachment; filename="+fileInfo.Name())

	// Set the Content-Type header based on the file's MIME type
	w.Header().Set("Content-Type", http.DetectContentType(nil))

	// Set the Content-Length header
	w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

	// Copy the file content to the response writer
	http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)
}

