package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Put(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "uploading file! \n")
	r.ParseMultipartForm(10 << 20) //10mb file max

	file, handler, err := r.FormFile("myFile")
	fileName := r.FormValue("fileName")
	if err != nil {
		fmt.Println("Error retrieveing file ")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded file: %v", handler.Filename)

	//create temp file in
	tempFile, err := ioutil.TempFile("temp-images/", fileName+"*.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	//copy contents
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)

	fmt.Fprintf(w, "Succesfully uploaded file \n")

}

func Get(rw http.ResponseWriter, req *http.Request) {
	fileName := req.FormValue("myFile")

    if fileName == "" {
        fmt.Fprint(rw, "Input not Recieved")
        fmt.Println("Input not Recieved")
        return
    }

	fileBytes, err := ioutil.ReadFile("temp-images/" + fileName)
	if err != nil {
        fmt.Fprint(rw, err.Error())
        fmt.Println(err.Error())
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/octet-stream")
	rw.Write(fileBytes)
}

func Delete(rw http.ResponseWriter, req *http.Request) {
	fileDel := req.FormValue("myFile")
	if fileDel == "" {
		fmt.Fprint(rw, "no input")
		fmt.Println("No input file")
		return
	}
	err := os.Remove("temp-images/" + fileDel)
	if err != nil {
		fmt.Fprint(rw, err.Error())
		fmt.Println(err.Error())
		return
	}
	fmt.Fprint(rw, "File deleted successfully!")
}

func List(rw http.ResponseWriter, req *http.Request) {
	files, err := ioutil.ReadDir("./temp-images/")
	if err != nil {
		fmt.Fprintf(rw, err.Error())
		return
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		fmt.Fprint(rw, file.Name(), " ", file.IsDir(), "\n")
	}

}

func Rename(rw http.ResponseWriter, req *http.Request) {
	baseDir := "./temp-images/"
	currName := req.FormValue("currName")
	newName := req.FormValue("newName")
	err := os.Rename(baseDir + currName, baseDir + newName)
	if err != nil {
		fmt.Fprintf(rw, currName)
		fmt.Fprintf(rw, newName)
		fmt.Fprintf(rw, err.Error())
		return
	}
	fmt.Fprint(rw, "File ", currName, " renamed to ", newName, " successfully!")
}
