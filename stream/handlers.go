package main

import (
	"crypto/md5"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func streamHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	videoId := params.ByName("video_id")
	videoL := VIDEO_DIR + videoId
	video, err := os.Open(videoL)
	if err != nil {
		sendErrorResponse(writer, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	writer.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(writer, request, "测试视频", time.Now(), video)
	defer video.Close()
}

func uploadHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	request.Body = http.MaxBytesReader(writer, request.Body, MAX_UPLOAD_SIZE)

	//校验文件大小
	if err := request.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(writer, http.StatusBadRequest, "File Is Too Big")
		return
	}
	// <form name file>
	//获取文件
	file, _, err := request.FormFile("file")
	if err != nil {
		log.Printf("Parse Request File Error : %v", err)
		sendErrorResponse(writer, http.StatusInternalServerError, "Parse Request File Error")
		return
	}
	//accept := "video/*"
	//读取二进制
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read File Error : %v", err)
		sendErrorResponse(writer, http.StatusInternalServerError, "Read File Error")
		return
	}
	fileName := params.ByName("video_id")
	//写文件
	err = ioutil.WriteFile(VIDEO_DIR+fileName, data, 0666)
	if err != nil {
		log.Printf("Write File Error : %v", err)
		sendErrorResponse(writer, http.StatusInternalServerError, "Write File Error")
		return
	}
	writer.WriteHeader(http.StatusCreated)
	io.WriteString(writer, "Upload Successfully")
}

func skipToUpLoadPageHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	currentTime := time.Now().Unix()
	hour := md5.New()
	io.WriteString(hour, strconv.FormatInt(currentTime, 10))

	//生成token
	token := fmt.Sprintf("%s", hour.Sum(nil))

	//渲染页面
	framework, _ := template.ParseFiles("./videos/upload.html")

	framework.Execute(writer, token)
}
