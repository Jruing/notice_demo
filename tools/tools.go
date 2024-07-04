package tools

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"
)

// 日志工具类
func InitLog(FileName string, MaxSize int, MaxAge int, MaxBackups int) *zap.Logger {
	/*
		FileName: 日志名称，包含日志所在路径
		MaxSize：每个日志文件的最大大小（MB）
		MaxAge: 保留旧文件的最大天数
		MaxBackups: 保留旧文件的最大数量
	*/
	if MaxSize == 0 {
		MaxSize = 500
	}
	if MaxAge == 0 {
		MaxAge = 30
	}
	if MaxBackups == 0 {
		MaxBackups = 14
	}
	loggerWriter := &lumberjack.Logger{
		Filename:   FileName,
		MaxSize:    MaxSize,
		MaxAge:     MaxAge,
		MaxBackups: MaxBackups,
		LocalTime:  false,
		Compress:   false,
	}
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:          "msg",
		LevelKey:            "level",
		TimeKey:             "ts",
		NameKey:             "logger",
		CallerKey:           "",
		FunctionKey:         "",
		StacktraceKey:       "",
		SkipLineEnding:      false,
		LineEnding:          zapcore.DefaultLineEnding,
		EncodeLevel:         zapcore.LowercaseLevelEncoder,
		EncodeTime:          zapcore.RFC3339TimeEncoder,
		EncodeDuration:      zapcore.StringDurationEncoder,
		EncodeCaller:        zapcore.ShortCallerEncoder,
		EncodeName:          nil,
		NewReflectedEncoder: nil,
		ConsoleSeparator:    "",
	}
	var (
		all_core []zapcore.Core
		core     zapcore.Core
	)
	all_core = append(all_core, zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(loggerWriter), // 使用 lumberjack 作为 zap 输出
		zap.NewAtomicLevelAt(zapcore.DebugLevel),
	))
	all_core = append(all_core, zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout), // 使用控制台作为zap输出
		zap.NewAtomicLevelAt(zapcore.DebugLevel),
	))

	core = zapcore.NewTee(all_core...)

	// 创建 Logger
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	defer logger.Sync() // 确保日志文件正常关闭

	// 使用 Logger 记录日志
	//logger.Info("This is an info level log", zap.String("key", "value"))
	return logger
}

// Request请求工具类
type LiteRequest struct {
	clients    http.Client
	urls       string
	methods    string
	headers    map[string]string
	jsondata   []byte
	formdata   string
	maxretries int
	returndata map[string]interface{}
	usehttp2   bool
	jar        cookiejar.Jar
}

func (receiver *LiteRequest) Session() *LiteRequest {
	// 创建一个 CookieJar 来存储 cookies
	jar, _ := cookiejar.New(nil)

	receiver.clients.Jar = jar
	return receiver
}

func (receiver *LiteRequest) AddHeader(headers map[string]string) *LiteRequest {
	receiver.headers = headers
	// 添加请求头
	return receiver
}

func (receiver *LiteRequest) AddProxy(proxy string) *LiteRequest {
	// 设置请求代理
	proxyurl, _ := url.Parse(proxy)
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyurl),
	}
	receiver.clients.Transport = transport
	return receiver
}

func (receiver *LiteRequest) AddTimeout(timeout time.Duration) *LiteRequest {
	receiver.clients.Timeout = timeout * time.Second
	return receiver
}

func (receiver *LiteRequest) SetRequestParams(urls string, methods string, jsondata []byte, formdata map[string]string) *LiteRequest {
	receiver.urls = urls
	switch methods {
	case "POST":
		fmt.Println("this is a POST request")
	case "GET":
		fmt.Println("this is a GET request")
	case "DELETE":
		fmt.Println("this is a DELETE request")
	case "PUT":
		fmt.Println("this is a PUT request")
	default:
		panic("error methods")
	}
	receiver.methods = methods
	receiver.jsondata = jsondata
	var r http.Request
	r.ParseForm()
	for key, value := range formdata {
		r.Form.Add(key, value)
	}
	bodystr := strings.TrimSpace(r.Form.Encode())
	receiver.formdata = bodystr
	return receiver
}

func (receiver *LiteRequest) AddRetryNum(maxretries int) *LiteRequest {
	// 设置最大重试次数
	receiver.maxretries = maxretries
	return receiver
}

func (receiver *LiteRequest) SendGetRequest() (map[string]interface{}, error) {
	var maxretries int
	if receiver.maxretries == 0 {
		maxretries = 1
	} else {
		maxretries = receiver.maxretries
	}
	for i := 0; i < maxretries; i++ {
		req, err := http.NewRequest(receiver.methods, receiver.urls, nil)
		if err != nil {
			panic(err)
		}
		for key, value := range receiver.headers {
			req.Header.Set(key, value)
		}
		response, err := receiver.clients.Do(req)
		defer response.Body.Close()
		if err == nil && response.StatusCode == http.StatusOK {
			returndata := make(map[string]interface{})
			fmt.Println(response.StatusCode)
			returndata["StatusCode"] = http.StatusOK
			body, _ := io.ReadAll(response.Body)
			returndata["Body"] = string(body)
			receiver.returndata = returndata
			return receiver.returndata, nil
		}
		fmt.Println("重试次数:", i+1)
		time.Sleep(3 * time.Second)
	}

	return receiver.returndata, errors.New("request failed after multiple attempts")
}

func (receiver *LiteRequest) SendJsonRequest() (map[string]interface{}, error) {
	var maxretries int
	if receiver.maxretries == 0 {
		maxretries = 1
	} else {
		maxretries = receiver.maxretries
	}
	for i := 0; i < maxretries; i++ {
		req, err := http.NewRequest(receiver.methods, receiver.urls, bytes.NewBuffer(receiver.jsondata))
		if err != nil {
			panic(err)
		}
		for key, value := range receiver.headers {
			req.Header.Set(key, value)
		}
		response, err := receiver.clients.Do(req)
		defer response.Body.Close()
		if err == nil && response.StatusCode == http.StatusOK {
			receiver.returndata["StatusCode"] = response.StatusCode
			body, _ := io.ReadAll(response.Body)
			receiver.returndata["Body"] = string(body)
			return receiver.returndata, nil
		}
		time.Sleep(3 * time.Second)
	}

	return receiver.returndata, errors.New("request failed after multiple attempts")
}

func (receiver *LiteRequest) SendFormRequest() (map[string]interface{}, error) {
	var maxretries int
	if receiver.maxretries == 0 {
		maxretries = 1
	} else {
		maxretries = receiver.maxretries
	}
	for i := 0; i < maxretries; i++ {
		req, err := http.NewRequest(receiver.methods, receiver.urls, strings.NewReader(receiver.formdata))
		if err != nil {
			panic(err)
		}
		for key, value := range receiver.headers {
			req.Header.Set(key, value)
		}
		response, err := receiver.clients.Do(req)
		defer response.Body.Close()
		if err == nil && response.StatusCode == http.StatusOK {
			receiver.returndata["StatusCode"] = response.StatusCode
			body, _ := io.ReadAll(response.Body)
			receiver.returndata["Body"] = string(body)
			return receiver.returndata, nil
		}
		time.Sleep(3 * time.Second)
	}

	return receiver.returndata, errors.New("request failed after multiple attempts")
}

func (receiver *LiteRequest) MapToJson(result map[string]interface{}) (bool, string) {
	jsonStringString, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Error marshalling mapStringString:", err)
	}
	isJson := json.Valid([]byte(fmt.Sprintf("%v", string(jsonStringString))))
	return isJson, string(jsonStringString)
}

// 加密
type EncryptoInterface interface {
	b64Encode(text string) string
	b64Decode(text string) string
	md5Encode(text string) string
}

type EncryptoStruct struct {
}

func (receiver EncryptoStruct) b64Encode(text string) string {
	enc_rs := base64.StdEncoding.EncodeToString([]byte(text))
	return enc_rs
}

func (receiver EncryptoStruct) b64Decode(text string) string {
	enc_rs, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		panic(fmt.Sprintf("解码异常,原始数据:%s", text))
	}
	return string(enc_rs)
}

func (receiver EncryptoStruct) md5Encode(text string) string {
	originalData := []byte(text)
	hash := md5.Sum(originalData)
	return string(hash[:])
}

// 生成随机字符串
func GenerateRandomString(length int) (string, error) {
	// 生成指定长度的安全随机字节
	b := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}

	// 将随机字节转换为字符串
	return fmt.Sprintf("%x", b), nil
}

// 返回值
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessResponse(msg string, data interface{}) *Response {
	return &Response{
		Code: 0, // 通常0表示成功
		Msg:  msg,
		Data: data,
	}
}

func ErrorResponse(errMsg string, code int) *Response {
	return &Response{
		Code: code,
		Msg:  errMsg,
	}
}
