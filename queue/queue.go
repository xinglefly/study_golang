package queue


//使用interface 代替所有类型
//可以在内部设置类型

type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v.(int))
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(string)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

//TODO go doc Queue  查看文档
// doc IsEmpty
//TODO godoc      生成文档
//godoc -help
//godoc -http :6060