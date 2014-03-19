package main

import (
  "./bst"
  "testing"
)

func ExampleBst() *bst.Tree {
  tree := &bst.Tree{}

  tree.Add(10)
  tree.Add(5)
  tree.Add(8)
  tree.Add(3)
  tree.Add(19)
  tree.Add(15)
  tree.Add(11)
  tree.Add(9)

  return tree
}

func TestBst(t *testing.T) {
  tree := ExampleBst()
  tree.Print()
}
