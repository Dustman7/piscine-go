package piscine

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	var left, right *TreeNode = root.Left, root.Right
	BTreeApplyPostorder(left, f)

	BTreeApplyPostorder(right, f)
	f(root.Data)
}
