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
      if (*parent).left == nil || value <= (*parent).left.value {
        parent = &((*parent).left)
      } else { // insert the element
        tmp := (*parent).left
        (*parent).left = &Node{value: value}
        (*parent).left.left = tmp
      }
    } else if value > (*parent).value {
      if (*parent).right == nil || value >= (*parent).right.value {
        parent = &((*parent).right)
      } else { // insert the element
        tmp := (*parent).right
        (*parent).right = &Node{value: value}
        (*parent).right.right = tmp
      }
    }
  }

  if *parent == nil {
    *parent = &Node{value: value}
  }
}

func (tree *Tree) Path (value int) {

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