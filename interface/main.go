package main

import (
	"fmt"
	"studygolang/interface/mock"
	"studygolang/interface/real"
	"os"
)

const url = "http://www.imooc.com"

type RetriverImp interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetriverPoster interface {
	RetriverImp
	Poster
}

func download(r RetriverImp) string {
	return r.Get("url")
}

func post(poster Poster) string {
	return poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

func session(retriver RetriverPoster) string {
	retriver.Post(url, map[string]string{
		"contents": "study golang forever!!!",
	})
	return retriver.Get(url)
}

func inspect(r RetriverImp) {
	fmt.Printf("%T %v\n", r, r)
	fmt.Println("Type switch")
	switch v := r.(type) {
	case *mock.Retriver:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriver:
		fmt.Println("UserAgent", v.UserAgent)
	}
}

func main() {

	var r RetriverImp
	r = &mock.Retriver{"implements two interface"}
	inspect(r)
	r = real.Retriver{"http1.2"}
	inspect(r)

	//fmt.Println(download(r))
	retriver := mock.Retriver{"this is fake contents."}

	fmt.Println("try to calling RetriverPoster:")
		fmt.Println(session(&retriver))

}
