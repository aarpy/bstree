package main

import (
  "./bst"
  "testing"
)

func ExampleBst() *bst.Tree {
  tree := &bst.Tree{}

  tree.Add(10)
  tree.Add(5)
  tree.Add(15)
  tree.Add(12)
  tree.Add(20)
  tree.Add(8)
  tree.Add(2)
  tree.Add(3)
  tree.Add(17)
  tree.Add(14)
  tree.Add(1)
  tree.Add(11)

  return tree
}

func TestBst(t *testing.T) {
  tree := ExampleBst()
  tree.Print()
}
