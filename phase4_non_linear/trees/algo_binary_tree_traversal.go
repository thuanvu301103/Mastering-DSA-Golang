package tree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// InOrder: Left -> Root -> Right
func InOrder(root *Node, result *[]int) {
	if root == nil {
		return
	}
	InOrder(root.Left, result)
	*result = append(*result, root.Value)
	InOrder(root.Right, result)
}

// PreOrder: Root -> Left -> Right
func PreOrder(root *Node, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Value)
	PreOrder(root.Left, result)
	PreOrder(root.Right, result)
}

// PostOrder: Left -> Right -> Root
func PostOrder(root *Node, result *[]int) {
	if root == nil {
		return
	}
	PostOrder(root.Left, result)
	PostOrder(root.Right, result)
	*result = append(*result, root.Value)
}
