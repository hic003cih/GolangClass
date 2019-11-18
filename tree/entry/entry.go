package main

import (
	"fmt"

	tree "../"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil {
		return
	}

	left :=myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder() 
	myTreeNode{myNode.node.Left}.postOrder()
	myTreeNode{myNode.node.Right}.postOrder()
	myNode.node.Print()

}

func main() {
	var root tree.TreeＮode

	root = TreeＮode{value: 3}
	root.Left = &TreeNode{}
	root.Right = &TreeNode{5, nil, nil}
	root.Right.Left = new(TreeNode)
	root.Left.Right = CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()
	fmt.Println()
	myTreeNode{&root}.postOrder()
	fmt.Println()
}
