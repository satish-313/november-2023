package code

import "fmt"

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

func findModeHM(root *treeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	m := make(map[int]int)

	dfs(root, m)

	mf := 0

	for _, f := range m {
		if f > mf {
			mf = f
		}
	}

	for i, v := range m {
		if v == mf {
			res = append(res, i)
		}
	}

	return res
}

func dfs(node *treeNode, m map[int]int) {
	if node == nil {
		return
	}

	m[node.val]++
	dfs(node.left, m)
	dfs(node.right, m)
}

func Test1() {
	c := treeNode{2, nil, nil}
	b := treeNode{2, &c, nil}
	a := treeNode{1, nil, &b}

	fmt.Println(findModeHM(&a))
}
