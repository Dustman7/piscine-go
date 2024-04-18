package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}

	data := root.Data
	if data == elem {
		return root
	}

	var left, right *TreeNode = root.Left, root.Right

	nodel := BTreeSearchItem(left, elem)
	if nodel != nil {
		return nodel
	}

	noder := BTreeSearchItem(right, elem)
	if noder != nil {
		return noder
	}

	return nil
}
