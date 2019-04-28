package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
		log.Printf("Read file error : %v", err)
		sendErrorResponse(writer, http.StatusInternalServerError, "服务异常")
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
