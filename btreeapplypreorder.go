package piscine

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	var left, right *TreeNode = root.Left, root.Right
	f(root.Data)
	BTreeApplyPreorder(left, f)

	BTreeApplyPreorder(right, f)
}
