package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const uploadDir = "./uploads"

func main() {
	os.MkdirAll(uploadDir, 0755)

	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/", fileHandler)
	http.HandleFunc("/delete/", deleteHandler)

	fmt.Println("Файловый сервер запущен на порту 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Ошибка получения файла", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filename := header.Filename
	filePath := filepath.Join(uploadDir, filename)

	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Ошибка сохранения файла", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Ошибка копирования файла", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Файл %s успешно загружен", filename)
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "Файловый сервер работает!")
		return
	}

	filename := r.URL.Path[1:]
	filePath := filepath.Join(uploadDir, filename)

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Файл не найден", http.StatusNotFound)
		return
	}
	defer file.Close()

	contentType := "application/octet-stream"
	switch filepath.Ext(filename) {
	case ".webp":
		contentType = "image/webp"
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	case ".gif":
		contentType = "image/gif"
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Cache-Control", "public, max-age=3600")

	io.Copy(w, file)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	filename := r.URL.Path[len("/delete/"):]
	filePath := filepath.Join(uploadDir, filename)

	err := os.Remove(filePath)
	if err != nil {
		http.Error(w, "Файл не найден", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Файл %s успешно удален", filename)
}
