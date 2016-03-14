package godragon

import "log"

type ItemTree struct {
	Nodes map[string]*ItemNode
}

func (it *ItemTree) Balance() {

}

func (it *ItemTree) AddNode(in *ItemNode) {
	//it.Nodes = append(it.Nodes, in)
}

type ItemNode struct {
	ID    string
	Child map[string]*ItemNode
}

func (in *ItemNode) AddChild(cn *ItemNode) {
	//in.Child = append(in.Child, cn)
}

func ItemBuildTree() error {

	items, err := StaticItems("6.4.1")
	if err != nil {
		return err
	}

	//it := ItemTree{Nodes: make(map[string]*ItemNode)}
	itemMap := make(map[string][]string)

	for i := range items {

		if len(items[i].Into) > 0 {

			for j := range items[i].Into {
				itemMap[items[i].ID] = append(itemMap[items[i].ID], items[i].Into[j])

			}
		}

	}

	log.Printf("%#v\n", itemMap)
	return nil
}
