package tree

import (
	"errors"
	"sort"
)

// Record represents a database record.
type Record struct {
	ID, Parent int
}

// Node represents a node in a tree.
type Node struct {
	ID       int
	Children []*Node
}

// Build constructs a tree of nodes from a set of records.
func Build(records []Record) (*Node, error) {
	if len(records) < 1 {
		return nil, nil
	}

	nodes := make(map[int]*Node) // to track seen nodes

	byID := func(i, j int) bool {
		return records[i].ID < records[j].ID
	}
	sort.Slice(records, byID)

	for i, rec := range records {
		var parent *Node

		if i == 0 {
			if rec.ID != 0 || rec.Parent != 0 {
				return nil, errors.New("invalid root")
			} 
		} else if i > 0 && (i != rec.ID || rec.ID <= rec.Parent) {
			return nil, errors.New("invalid record")
		}

		/*
		// check various tree conditions
		if i != rec.ID {
			return nil, errors.New("sequence is non-continuous")
		} else if i > 0 {
			var ok bool
			if parent, ok = nodes[rec.Parent]; !ok {
				return nil, errors.New("parent missing")
			}
		} else if rec.Parent != 0 {
			return nil, errors.New("root cannot have parent")
		}
		*/

		nodes[i] = &Node{ID: i}

		// if parent != nil {
		if parent, ok = nodes[rec.Parent]; !ok {
			parent.Children = append(parent.Children, nodes[i])
		}
	}

	return nodes[0], nil
}
