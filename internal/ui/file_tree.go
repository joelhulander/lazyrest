package ui

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var fileTree *FileTree 

type FileTree struct {
	tree    *tview.TreeView
	root    *tview.TreeNode
	rootDir string
	errorLogger *log.Logger
}

type treeNode struct {
	path  string
	isDir bool
	name  string
}

func NewFileTree(rootDir string, errorLogger *log.Logger) *FileTree {
	root := tview.NewTreeNode(rootDir).SetColor(tcell.ColorRed)

	tree := tview.NewTreeView()
	tree.
		SetTopLevel(1).
		SetGraphics(false).
		SetRoot(root).
		SetCurrentNode(root).
		SetBorder(true)

	tree.SetTitle("Collections")

	ft := &FileTree{
		tree:    tree,
		rootDir: rootDir,
		root: root,
		errorLogger: errorLogger,
	}

	err := ft.addChildren(ft.root, ft.rootDir)
	
	if err != nil {
		ft.errorLogger.Println(err)
	}

	ft.tree.SetSelectedFunc(ft.handleSelected)

	fileTree = ft

	return ft
}

func (ft *FileTree) addChildren(target *tview.TreeNode, path string) error {
	treeItems, err := os.ReadDir(path)
	if err != nil {
		return err
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

	return nil
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

	node.SetSelectedFunc(func() { ft.fileSelected(reference.path) })
}

func (ft *FileTree) GetView() tview.Primitive {
	return ft.tree
}

func (ft *FileTree) fileSelected(path string) {
	fileContent, err := os.ReadFile(path)

	if err != nil {
		ft.errorLogger.Println(err)
	}

	responseView.textArea.SetText(string(fileContent))
}
