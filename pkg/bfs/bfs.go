package bfs

type Establishment struct {
	Verified   bool
	Name       string
	IsKeychain bool
}

type GraphItem struct {
	Value     *Establishment
	Neighbors []string
}

type Graph map[string]GraphItem

func NewEstablishment(name string, isKeychain bool) *Establishment {
	e := Establishment{
		Name:       name,
		IsKeychain: isKeychain,
	}

	return &e
}

func FindKeychain(stablishments Graph, rootName string) *Establishment {
	q := NewQueue[string]()

	root, ok := stablishments[rootName]

	if !ok {
		return nil
	}

	if root.Value.IsKeychain {
		return root.Value
	} else {
		root.Value.Verified = true
	}

	for q.EnqueueMany(root.Neighbors); q.Head != nil; {
		nodeName := q.Dequeue()

		node := stablishments[nodeName.Value]

		if node.Value.Verified {
			continue
		}

		node.Value.Verified = true

		if node.Value.IsKeychain {
			return node.Value
		}

		q.EnqueueMany(node.Neighbors)
	}

	return nil
}
