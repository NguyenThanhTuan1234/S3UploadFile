package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/NguyenThanhTuan1234/S3UploadFile/config"
	"github.com/NguyenThanhTuan1234/S3UploadFile/s3client"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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
		profile := flag.String("profile", "", "Enter AWS Profile")
		bucket := flag.String("bucket", "", "Enter bucket name")
		flag.Parse()
		awsConfigProvider := config.AWSConfigProvider{
			Profile: *profile,
		}
		awsConfig, err := awsConfigProvider.LoadAWSConfig()
		if err != nil {
			fmt.Println(err)
		}
		putObject := s3client.S3New(awsConfig)
		err = putObject.Upload(buffer, nf.Name(), *bucket)
		if err != nil {
			fmt.Println(err)
		}
	}
	tpl.ExecuteTemplate(w, "index.html", nil)
}
