type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	var helper func(l, r int) *TreeNode

	helper = func(l, r int) *TreeNode {
		// se não tiver elementos no subarray atual vamos retornar nil
		if l > r {
			return nil
		}

		m := (l + r) / 2
		root := &TreeNode{Val: nums[m]}
		root.Left = helper(l, m-1)
		root.Right = helper(m+1, r)

		return root
	}

	return helper(0, len(nums)-1)
}