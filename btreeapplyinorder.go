package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	var left, right *TreeNode = root.Left, root.Right
	BTreeApplyInorder(left, f)
	f(root.Data)
	BTreeApplyInorder(right, f)
}

/* fmt.Println(root.Left.Data)
	fmt.Println(root.Data)
	fmt.Println(root.Right.Left.Data)
	fmt.Println(root.Right.Data)
} */
