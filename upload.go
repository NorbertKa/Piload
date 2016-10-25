package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/nu7hatch/gouuid"
)

var (
	ErrCouldntUploadFile   = errors.New("Couldn't upload file")
	ErrUnsupportedFileType = errors.New("Unsupported file type")
)

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(404)
			w.Write([]byte(ErrCouldntUploadFile.Error()))
			return
		}
		contentType := handler.Header.Get("Content-Type")
		fmt.Println(contentType)
		if contentType == "image/jpeg" || contentType == "image/png" || contentType == "image/gif" || contentType == "image/svg+xml" {
			defer file.Close()
			ID, err := uuid.NewV4()
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(404)
				w.Write([]byte(ErrCouldntUploadFile.Error()))
				return
			}
			f, err := os.OpenFile("./static/"+ID.String(), os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(404)
				w.Write([]byte(ErrCouldntUploadFile.Error()))
				return
			}
			defer f.Close()
			io.Copy(f, file)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(404)
				w.Write([]byte(ErrCouldntUploadFile.Error()))
				return
			}
			http.Redirect(w, r, "/static/"+ID.String(), http.StatusSeeOther)
		} else {
			fmt.Println(err)
			w.WriteHeader(404)
			w.Write([]byte(ErrUnsupportedFileType.Error()))
			return
		}
	}
}
