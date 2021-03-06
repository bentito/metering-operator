// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/kubernetes-reporting/metering-operator/pkg/apis/metering/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HiveTableLister helps list HiveTables.
type HiveTableLister interface {
	// List lists all HiveTables in the indexer.
	List(selector labels.Selector) (ret []*v1.HiveTable, err error)
	// HiveTables returns an object that can list and get HiveTables.
	HiveTables(namespace string) HiveTableNamespaceLister
	HiveTableListerExpansion
}

// hiveTableLister implements the HiveTableLister interface.
type hiveTableLister struct {
	indexer cache.Indexer
}

// NewHiveTableLister returns a new HiveTableLister.
func NewHiveTableLister(indexer cache.Indexer) HiveTableLister {
	return &hiveTableLister{indexer: indexer}
}

// List lists all HiveTables in the indexer.
func (s *hiveTableLister) List(selector labels.Selector) (ret []*v1.HiveTable, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.HiveTable))
	})
	return ret, err
}

// HiveTables returns an object that can list and get HiveTables.
func (s *hiveTableLister) HiveTables(namespace string) HiveTableNamespaceLister {
	return hiveTableNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HiveTableNamespaceLister helps list and get HiveTables.
type HiveTableNamespaceLister interface {
	// List lists all HiveTables in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.HiveTable, err error)
	// Get retrieves the HiveTable from the indexer for a given namespace and name.
	Get(name string) (*v1.HiveTable, error)
	HiveTableNamespaceListerExpansion
}

// hiveTableNamespaceLister implements the HiveTableNamespaceLister
// interface.
type hiveTableNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HiveTables in the indexer for a given namespace.
func (s hiveTableNamespaceLister) List(selector labels.Selector) (ret []*v1.HiveTable, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.HiveTable))
	})
	return ret, err
}

// Get retrieves the HiveTable from the indexer for a given namespace and name.
func (s hiveTableNamespaceLister) Get(name string) (*v1.HiveTable, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("hivetable"), name)
	}
	return obj.(*v1.HiveTable), nil
}
