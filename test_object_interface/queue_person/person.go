package queue_person


type Queue []interface{}

func (q *Queue) Push(obj interface{}) {
	*q = append(*q, obj)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) Traverse() interface{} {
	return (*q)[:]
}

func (q *Queue) Peek() interface{} {
	return (*q)[0]
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}
