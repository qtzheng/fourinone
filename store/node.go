package store

type node struct {
	Key   string
	Value string
	store *store
}

func newNode(store *store, key, value string) *node {
	return &node{
		Key:   key,
		store: store,
		Value: value,
	}
}
