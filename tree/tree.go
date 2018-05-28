package tree

import "fmt"

type Node struct {
	Left, Right *Node
	Value       int
}

//工厂模式
func CreateNode(value int) *Node {
	return &Node{Value: value} //创建在堆上还是栈上,go语言无所谓
}

func (node Node) Print() {
	fmt.Print(node.Value," ")
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("node is nil isIgnore!")
		return
	}
	node.Value = value
}

