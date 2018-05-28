package mock

import "fmt"

type Retriver struct {
	Contents string
}

func (r *Retriver) String() string {
	return fmt.Sprintf("Retriever: {contents = %s}", r.Contents)
}

func (r *Retriver) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

func (r *Retriver) Get(url string) string {
	return r.Contents
}

//函数是一等公民：参数，变量，返回值都可以是函数
//高阶函数
//