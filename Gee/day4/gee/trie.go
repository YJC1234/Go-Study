package gee

import "strings"

//通过前缀树表示路由，可以实现动态路由
//:匹配part，*匹配之后所有部分

type node struct {
	pattern  string  //待匹配路由，如/p/:lang/,中途没有
	part     string  //路由中这一节点的部分
	children []*node //子节点
	isWild   bool    //精确匹配？当part以*或:开头为true
}

// -------------辅助方法-----------
// 从n的孩子中找出第一个匹配part的节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 从n的孩子中找出所有匹配part的节点
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

//--------------------------------

// 插入节点到n树中. pattern为全路由，parts为路由切割后，height为高度(一般传入0，递归专用)
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == '*' || part[0] == ':'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// 查找n树中符合parts的节点
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}

// 获取n树所有pattern
func (n *node) travel(list *([]*node)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}
