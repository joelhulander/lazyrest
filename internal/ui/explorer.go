package ui

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var explorer *CollectionsExplorer 

type CollectionsExplorer struct {
	view    *tview.TreeView
	root    *tview.TreeNode
	rootDir string
	logger *slog.Logger
}

type treeNode struct {
	path  string
	isDir bool
	name  string
}

func NewCollectionsExplorer(rootDir string, logger *slog.Logger) *CollectionsExplorer {
	root := tview.NewTreeNode(rootDir)

	treeView := tview.NewTreeView()
	treeView.
		SetTopLevel(1).
		SetGraphics(false).
		SetRoot(root).
		SetCurrentNode(root).
		SetBorder(true)

	treeView.SetTitle(" [1] Collections ").SetTitleAlign(0).SetBorderPadding(1,0,1,0)

	ft := &CollectionsExplorer{
		view:    treeView,
		rootDir: rootDir,
		root: root,
		logger: logger,
	}

	err := ft.addChildren(ft.root, ft.rootDir)
	if err != nil {
		logger.Error("error occurred", "err", err)
	}

	ft.view.SetSelectedFunc(ft.handleSelected)

	explorer = ft

	return ft
}

func (ft *CollectionsExplorer) addChildren(target *tview.TreeNode, path string) error {
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

func (ft *CollectionsExplorer) handleSelected(node *tview.TreeNode) {
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

func (ft *CollectionsExplorer) GetView() *tview.TreeView {
	return ft.view
}

func (ft *CollectionsExplorer) fileSelected(path string) {
	fileContent, err := os.ReadFile(path)

	if err != nil {
		ft.logger.Error("error occurred", "err", err)
	}

	responseView.view.SetText(string(fileContent))
}
