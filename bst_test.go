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
  tree.Add(13)

  return tree
}

func TestBstInsertion(t *testing.T) {
  tree := ExampleBst()
  tree.Print()
}

func TestBstPath(t *testing.T) {
  tree := ExampleBst()
  tree.Path(3)
  tree.Path(8)
  tree.Path(1)
  tree.Path(17)
  tree.Path(13)
  tree.Path(16)
}
