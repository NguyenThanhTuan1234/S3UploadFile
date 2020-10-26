package main

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

var tpl *template.Template
var svcS3 *s3.Client

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		mf, fh, err := r.FormFile("nf")
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close()
		nf, err := os.Create(fh.Filename)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()
		mf.Seek(0, 0)
		io.Copy(nf, mf)
		fileInfo, _ := nf.Stat()
		size := fileInfo.Size()
		buffer := make([]byte, size)
		fileBytes := bytes.NewReader(buffer)
		cfg, err := config.LoadDefaultConfig()
		if err != nil {
			panic(err)
		}
		svcS3 = s3.NewFromConfig(cfg)
		putObjectInput := s3.PutObjectInput{
			Bucket: aws.String("s3testtuan1"),
			Key:    aws.String(nf.Name()),
			Body:   fileBytes,
		}
		_, err = svcS3.PutObject(context.Background(), &putObjectInput)
		if err != nil {
			fmt.Println(err)
		}
	}
	tpl.ExecuteTemplate(w, "index.html", nil)
}
