package _35_lowestCommonAncestor

import (
	"testing"
)

//         6
//       /   \
//      2      8
//     /  \   /  \
//     0   4  7   9
//        / \
//       3   5
func TestLowestCommonAncestor(t *testing.T) {
	node := new(TreeNode)
	node.Val = 6
	node.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}}
	node.Right = &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}}
	ret := LowestCommonAncestor(node, node.Left, node.Right)
	t.Log(ret.Val)

	ret2 := LowestCommonAncestor(node, node.Left, &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}})
	t.Log(ret2.Val)

	ret3 := LowestCommonAncestor(node, node.Left.Left, &TreeNode{Val: 5})
	t.Log(ret3.Val)

	t.Log("=====================================")

	ret4 := LowestCommonAncestor2(node, node.Left, node.Right)
	t.Log(ret4.Val)

	ret5 := LowestCommonAncestor2(node, node.Left, &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}})
	t.Log(ret5.Val)

	ret6 := LowestCommonAncestor2(node, node.Left.Left, &TreeNode{Val: 5})
	t.Log(ret6.Val)
}
