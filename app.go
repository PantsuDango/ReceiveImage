package main

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Params struct {
	Language    string `json:"Language" binding:"required"`
	ImageBase64 string `json:"ImageBase64" binding:"required"`
}

// 初始化
func init() {

	// 日志参数
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// 初始化图片保存目录
	CreateDateDir("images")
	for _, val := range []string{"ENG", "JAP", "KOR"} {
		CreateDateDir(filepath.Join("images", val))
	}
}

// 成功的返回
func JSONSuccess(ctx *gin.Context, result interface{}) {

	ctx.JSON(http.StatusOK, gin.H{
		"Code":      0,
		"Status":    "Success",
		"RequestID": uuid.NewV4(),
		"Result":    result,
	})
}

// 失败的返回
func JSONFail(ctx *gin.Context, errorMsg string) {

	ctx.JSON(http.StatusOK, gin.H{
		"Code":      -1,
		"Status":    "Fail",
		"RequestID": uuid.NewV4(),
		"ErrorMsg":  errorMsg,
	})
}

// 创建目录
func CreateDateDir(folderPath string) {

	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		_ = os.Mkdir(folderPath, 0777)
	}
}

// 收图接口
func ReceiveImg(ctx *gin.Context) {

	// 校验请求参数
	var params Params
	if err := ctx.ShouldBindBodyWith(&params, binding.JSON); err != nil {
		JSONFail(ctx, err.Error())
		log.Println("Request Params Error: ", err)
		return
	}
	// 校验Language
	switch params.Language {
	case "ENG", "JAP", "KOR":
	default:
		JSONFail(ctx, "Language error")
		log.Println("Request Params Error: ", "Language error")
		return
	}
	// 图片base64解码
	imgBytes, err := base64.StdEncoding.DecodeString(params.ImageBase64)
	if err != nil {
		JSONFail(ctx, err.Error())
		log.Println("Image Decode Error: ", err)
		return
	}
	// 按照日期创建文件夹
	folderPath := filepath.Join("images", params.Language, time.Now().Format("2006-01-02"))
	CreateDateDir(folderPath)
	// 保存图片
	now := strings.Replace(time.Now().Format("20060102-150405.999999999"), ".", "-", -1)
	fileName := fmt.Sprintf(`%s_%s.jpg`, params.Language, now)
	filePath := filepath.Join(folderPath, fileName)
	err = ioutil.WriteFile(filePath, imgBytes, 0666)
	if err != nil {
		JSONFail(ctx, err.Error())
		log.Println("Image Save Error: ", err)
		return
	}
	// 成功返回
	log.Println("Save Image Success: ", filePath)
	JSONSuccess(ctx, "Success")
}

// 从命令行参数获取端口号
func GetPort() string {

	var port string
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	port = os.Args[1]

	return port
}

func main() {

	// 从命令行获取端口号, 不传默认3000
	port := GetPort()
	if port == "" {
		port = "3000"
	}

	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/receive_img", ReceiveImg)

	err := router.Run(":" + port)
	if err != nil {
		log.Println("Server run error: ", err)
	}
}
