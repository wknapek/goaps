package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Records []Record

type Node struct {
	ID       int
	Children []*Node
}

func Build(rec Records) (*Node, error) {
	sort.Slice(rec, func(i int, j int) bool { return rec[i].ID < rec[j].ID })
	tree := map[int]*Node{}
	for idx, re := range rec {
		if re.ID != idx || re.ID < re.Parent || re.ID > 0 && re.Parent == re.ID {
			return nil, errors.New("not in sequence or bad parent")
		}
		tree[re.ID] = &Node{ID: re.ID}
		if idx > 0 {
			p := tree[re.Parent]
			p.Children = append(p.Children, tree[re.ID])
		}
	}
	return tree[0], nil
}
