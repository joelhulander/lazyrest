package ui

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/joelhulander/lazyrest/internal/appctx"
	"github.com/rivo/tview"
)

var explorer *CollectionsExplorer 

type CollectionsExplorer struct {
	ctx *appctx.Context
	view    *tview.TreeView
	root    *tview.TreeNode
	rootDir string
}

type treeNode struct {
	path  string
	isDir bool
	name  string
}

func NewCollectionsExplorer(ctx *appctx.Context, rootDir string) *CollectionsExplorer {
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
		ctx: ctx,
		view:    treeView,
		rootDir: rootDir,
		root: root,
	}

	_, err := ft.addChildren(ft.root, ft.rootDir)
	if err != nil {
		log.Error("error occurred", "err", err)
	}

	ft.view.SetSelectedFunc(ft.handleSelected)

	explorer = ft

	return ft
}

func (ft *CollectionsExplorer) addChildren(target *tview.TreeNode, path string) ([]string, error) {
	treeItems, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var nodes []string

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
		nodes = append(nodes, node.GetText())
	}

	return nodes, nil
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
			toggleNode(node, reference.name)
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

	node.SetSelectedFunc(func() {
		ft.ctx.OnFileSelected(reference.path)
	})
}

func toggleNode(node *tview.TreeNode, name string) {
	if node.GetText() == " ▼ " + name {
		node.SetText(" ▶ " + name)
		return 
	}
	node.SetText(" ▼ " + name)
}

func (ft *CollectionsExplorer) GetView() *tview.TreeView {
	return ft.view
}

