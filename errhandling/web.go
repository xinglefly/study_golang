package main

import (
	"net/http"
	"studygolang/errhandling/filelistingserver/filelisting"
	"log"
	"os"
)

//意料之中的事情使用error
//意料之外的： 使用panic.如数组越界\

//表格驱动测试


//两边的都是定义各自的error
type userError interface {
	error
	Message() string
}

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic :%v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)

		if err != nil {
			//logs.Warn("Error handling request: %s", err.Error())
			log.Printf("Error handling request: %s", err.Error())
			if userError, ok := err.(userError); ok {
				http.Error(writer,
					userError.Message(),
					http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err): //TODO 主要目的是作为对外展示不显示程序内部错误
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,
				http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandlerFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

//TODO http测试  1.通过使用假的Request/Response（相当于堆代码为了测试所有的可能性，速度快，粒度更细）  2.通过起服务器，测试服务器的性能（支撑度比较高，代码的覆盖量大）