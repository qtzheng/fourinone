package store

import (
	err "github.com/qtzheng/fourinone/error"
	"sync"
)

const (
	defaultVersion = 2
)

type Store interface {
	Version() uint64
	GetDomain(domain string) (*domain, error)
	GetNode(domain, node string) (*node, error)
	DeleteDomain(domain string)
	DeleteNode(domain, node string)
	PutNode(domain, node, newValue string) error
}

func New() Store {
	s := new(store)
	s.version = defaultVersion
	s.data = make(map[string]*domain)
	return s
}

type domain struct {
	lock  sync.RWMutex
	data  map[string]*node
	store *store
}

func newDomain(store *store) *domain {
	d := &domain{
		data:  make(map[string]*node),
		store: store,
	}
	return d
}
func (d *domain) getNode(node string) (*node, error) {
	d.lock.RLock()
	defer d.lock.RUnlock()
	if n, ok := d.data[node]; ok {
		return n, nil
	} else {
		return nil, err.NewError(err.KeyNotFound, "找不到该节点", nil)
	}
}
func (d *domain) deleteNode(node string) {
	d.lock.Lock()
	defer d.lock.Unlock()
	delete(d.data, node)
}
func (d *domain) putNode(node, value string) {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.data[node] = newNode(d.store, node, value)
}

type store struct {
	data    map[string]*domain
	lock    sync.RWMutex
	version uint64
}

func (s *store) Version() uint64 {
	return s.version
}
func (s *store) GetDomain(domain string) (*domain, error) {
	return s.getDomain(domain)
}
func (s *store) getDomain(domain string) (*domain, *err.Error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	d, ok := s.data[domain]
	if !ok {
		return nil, err.NewError(err.KeyNotFound, "找不到该domain", nil)
	}
	return d, nil
}

func (s *store) GetNode(domain, node string) (*node, error) {
	d, err := s.getDomain(domain)
	if err != nil {
		return nil, err
	}
	return d.getNode(node)
}
func (s *store) DeleteDomain(domain string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.data, domain)
}
func (s *store) DeleteNode(domain, node string) {
	d, err := s.GetDomain(domain)
	if err != nil {
		return
	}
	d.deleteNode(node)
}
func (s *store) PutNode(domain, node, newValue string) error {
	d, e := s.getDomain(domain)
	if d != nil {
		if e.ErrorCode == err.KeyNotFound {
			d = newDomain(s)
		} else {
			return e
		}
	}
	d.putNode(node, newValue)
	return nil
}
