package web

import "strings"

type node struct {
	pattern  string  // 待匹配路由
	part     string  // 路由中的一部分
	children []*node // 子节点,例如 [doc, tutorial, intro]
	isWild   bool    // 是否是模糊匹配
}

// matchChild 第一个匹配的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// matchChildren 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert 插入
func (n *node) insert(pattern string, parts []string, height int) {
	// 已经到达尽头，把pattern记录下来
	if len(parts) == height {
		// 末端树记录
		n.pattern = pattern
		return
	}
	// 非末端，将当前节点记录下来/p/d/doc, 例如doc
	part := parts[height]
	// 找一下存不存在该路径
	child := n.matchChild(part)
	// 如果不存在就新建一个
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	// 到达末端,或者碰到任意通配符
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	// 传入路径某一层级，例如go
	part := parts[height]
	// 在字典树中找满足的层级，例如精确匹配的或者通配的
	children := n.matchChildren(part)
	// 每一个层级下去查找
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
