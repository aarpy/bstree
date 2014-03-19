package bst

import (
  "fmt"
  "strconv"
)

type Node struct {
  value int
  right *Node
  left  *Node
}

type Tree struct {
  root  *Node
}

func (tree *Tree) Add (value int) {

  parent := &tree.root

  for *parent != nil {
    if value == (*parent).value {
      break //do not worry about duplicates
    } else if value < (*parent).value {
      parent = &((*parent).left)
    } else {
      parent = &((*parent).right)
    }
  }

  if *parent == nil {
    *parent = &Node{value: value}
  }
}

func (tree *Tree) Path (value int) {

  found := false
  node := tree.root

  for node != nil {
    if value == node.value {
      fmt.Println(value)
      found = true
      break;
    } else if (value < node.value) {
      fmt.Print(node.value, " L> ")
      node = node.left
    } else {
      fmt.Print(node.value, " R> ")
      node = node.right
    }
  }

  if !found {
    fmt.Println(value, "- Not found")
  }
}

func (tree *Tree) Print () {
  fmt.Println("Tree structure")
  if tree.root != nil {
    tree.root.print()
  } else {
    fmt.Println(" * empty tree * ")
  }
}

func (node *Node) print() {
  if node == nil {
    return
  }

  left := "nil"
  if node.left != nil {
    left = strconv.Itoa(node.left.value)
  }

  right := "nil"
  if node.right != nil {
    right = strconv.Itoa(node.right.value)
  }

  fmt.Println(" ", node.value, ":", left, ",", right)

  if node.left != nil {
    node.left.print()
  }

  if node.right != nil {
    node.right.print()
  }
}