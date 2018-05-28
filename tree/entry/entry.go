package main

import (
	"studygolang/tree"
	"fmt"
)

//组合的方式，扩展方式的写法，不修改原有的代码
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {

	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{nil, nil, 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Right.Left.SetValue(4)
	/*root.print()

	root.Right.Left.print()

	var pRoot *Node
	pRoot.setValue(200)
	pRoot = &root
	pRoot.setValue(300)
	pRoot.print()*/
	fmt.Println("前续遍历：")
	root.Traverse()
	fmt.Println("\n后续遍历：")
	node := myTreeNode{&root}
	node.postOrder()

	fmt.Println("\n函数式调用：   ")
	root.Traversef()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node Count:", nodeCount)

	c := root.TraverseWithChannel()
	maxNum := 0
	for node := range c {
		if node.Value > maxNum {
			maxNum = node.Value
		}
	}
	fmt.Println("最大节点数：", maxNum)

}
