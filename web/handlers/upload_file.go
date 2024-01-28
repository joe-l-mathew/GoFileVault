package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joe-l-mathew/GoFileVault/models"
	"github.com/joe-l-mathew/GoFileVault/pkg"
	"github.com/joe-l-mathew/GoFileVault/pkg/db"
	"github.com/joe-l-mathew/validogo"
)

func UploadFiles(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userId := ctx.Value("userId")

	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	// Get the file from the request
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error getting file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()
	filePath := filepath.Join(pkg.StorageName, fmt.Sprint(userId), ".", handler.Filename)
	outputFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()
	fileType := validogo.DetectFileType(handler.Filename)
	// Copy the file content to the newly created file
	_, err = io.Copy(outputFile, file)
	var fileModel models.UserFiles = models.UserFiles{
		UserId:   fmt.Sprint(userId),
		Filetype: string(fileType),
		FileName: handler.Filename,
		FilePath: filePath,
	}
	if err := db.DB.Create(&fileModel).Error; err != nil {
		http.Error(w, "Error adding file to db", http.StatusInternalServerError)
		return
	}
	if err != nil {
		http.Error(w, "Error copying file content", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("File uploaded successfully"))

}
