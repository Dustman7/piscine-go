package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	parent := node.Parent
	if node == parent.Left {
		parent.Left = rplc
		rplc.Parent = rplc
	} else if node == parent.Right {
		parent.Right = rplc
		rplc.Parent = parent
	}
	return root
}
