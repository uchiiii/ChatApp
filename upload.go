package main

import (
	"io/ioutil"
	"net/http"
	"io"
	"path"
)

func uploaderHandler(w http.ResponseWriter, req *http.Request) {
	userId := req.FormValue("userid")
	file, header, err := req.FormFile("avatarFile") //just return the file itself with the multipart.File interface type
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) //http.StatusInternalServerError is 500 error.
		return 
	}
	data, err := ioutil.ReadAll(file) //keep reading from the specified io.Reader interface until all of the bytes have been recieved.
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	filename := path.Join("avatars", userId+path.Ext(header.Filename)) //path.Ext return extension of the file.
	err = ioutil.WriteFile(filename, data, 0777) //0777 is permission.
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, "Successful")
}