package tree

import "fmt"

type TreeＮode struct {
	Ｖalue
	Left, Right *TreeＮode
}

func (node TreeＮode) Print() {
	fmt.Print(node.value, "")
}

func (node *TreeＮode) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil" + "node. Ignored.")
		return
	}
	node.Value = value
}

func CreateNode(value int) *TreeＮode {
	return &TreeＮode(Value: value)
}