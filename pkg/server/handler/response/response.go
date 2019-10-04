package response

import (
	"bytes"
	"encoding/json"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"
)

// Success HTTPコード:200 正常終了を処理する
func Success(writer http.ResponseWriter, response interface{}) {
	data, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		InternalServerError(writer, "marshal error")
		return
	}
	writer.Write(data)
}

func SuccessCode(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusOK)
}

// BadRequest HTTPコード:400 BadRequestを処理する
func BadRequest(writer http.ResponseWriter, message string) {
	httpError(writer, http.StatusBadRequest, message)
}

// StatusUnauthorized HTTPコード:401 StatusUnauthorizedを処理する
func StatusUnauthorized(writer http.ResponseWriter, message string) {
	httpError(writer, http.StatusUnauthorized, message)
}

// InternalServerError HTTPコード:500 InternalServerErrorを処理する
func InternalServerError(writer http.ResponseWriter, message string) {
	httpError(writer, http.StatusInternalServerError, message)
}

// httpError エラー用のレスポンス出力を行う
func httpError(writer http.ResponseWriter, code int, message string) {
	data, _ := json.Marshal(errorResponse{
		Code:    code,
		Message: message,
	})
	writer.WriteHeader(code)
	if data != nil {
		writer.Write(data)
	}
}

func SuccessFile(writer http.ResponseWriter, img *image.Image) {
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("unable to encode image.")
	}
	writer.Header().Set("Content-Type", "image/jpeg")
	writer.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := writer.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

// errorResponse エラー時の返り値
type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
