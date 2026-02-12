package ui

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type FileTree struct {
	tree    *tview.TreeView
	root    *tview.TreeNode
	rootDir string
}

type treeNode struct {
	path  string
	isDir bool
	name  string
}

func NewFileTree(rootDir string) *FileTree {
	ft := &FileTree{
		tree:    tview.NewTreeView().SetTopLevel(1).SetGraphics(false),
		rootDir: rootDir,
		root:    tview.NewTreeNode(rootDir).SetColor(tcell.ColorRed),
	}

	ft.addChildren(ft.root, ft.rootDir)

	ft.tree.
		SetRoot(ft.root).
		SetCurrentNode(ft.root).
		SetSelectedFunc(ft.handleSelected)

	return ft
}

func (ft *FileTree) addChildren(target *tview.TreeNode, path string) {
	treeItems, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, item := range treeItems {
		reference := treeNode{filepath.Join(path, item.Name()), item.IsDir(), item.Name()}
		node := tview.NewTreeNode("").SetReference(reference)

		if item.IsDir() {
			node.SetText(" ▶ " + item.Name())
			node.SetColor(tcell.ColorRed)
		} else {
			itemName := strings.TrimSuffix(item.Name(), filepath.Ext(item.Name()))
			node.SetText(itemName)
		}

		target.AddChild(node)
	}
}

func (ft *FileTree) handleSelected(node *tview.TreeNode) {
	if node.GetReference() == nil {
		return
	}

	reference := node.GetReference().(treeNode)

	if reference.isDir {
		children := node.GetChildren()

		if len(children) == 0 {
			path := reference.path
			ft.addChildren(node, path)
			node.SetText(" ▼ " + reference.name)
			return
		}

		if node.IsExpanded() {
			node.SetExpanded(false)
			node.SetText(" ▶ " + reference.name)
			return
		}

		node.SetExpanded(true)
		node.SetText(" ▼ " + reference.name)

		return
	}

	node.SetSelectedFunc(func() { fileSelected(reference.path) })
}

func fileSelected(path string) {
	_, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}
}

func (ft *FileTree) GetView() *tview.TreeView {
	return ft.tree
}
