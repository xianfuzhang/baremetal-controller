/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "./pkg/apis/baremetal/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BaremetalLister helps list Baremetals.
// All objects returned here must be treated as read-only.
type BaremetalLister interface {
	// List lists all Baremetals in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Baremetal, err error)
	// Baremetals returns an object that can list and get Baremetals.
	Baremetals(namespace string) BaremetalNamespaceLister
	BaremetalListerExpansion
}

// baremetalLister implements the BaremetalLister interface.
type baremetalLister struct {
	indexer cache.Indexer
}

// NewBaremetalLister returns a new BaremetalLister.
func NewBaremetalLister(indexer cache.Indexer) BaremetalLister {
	return &baremetalLister{indexer: indexer}
}

// List lists all Baremetals in the indexer.
func (s *baremetalLister) List(selector labels.Selector) (ret []*v1alpha1.Baremetal, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Baremetal))
	})
	return ret, err
}

// Baremetals returns an object that can list and get Baremetals.
func (s *baremetalLister) Baremetals(namespace string) BaremetalNamespaceLister {
	return baremetalNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BaremetalNamespaceLister helps list and get Baremetals.
// All objects returned here must be treated as read-only.
type BaremetalNamespaceLister interface {
	// List lists all Baremetals in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Baremetal, err error)
	// Get retrieves the Baremetal from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Baremetal, error)
	BaremetalNamespaceListerExpansion
}

// baremetalNamespaceLister implements the BaremetalNamespaceLister
// interface.
type baremetalNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Baremetals in the indexer for a given namespace.
func (s baremetalNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Baremetal, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Baremetal))
	})
	return ret, err
}

// Get retrieves the Baremetal from the indexer for a given namespace and name.
func (s baremetalNamespaceLister) Get(name string) (*v1alpha1.Baremetal, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("baremetal"), name)
	}
	return obj.(*v1alpha1.Baremetal), nil
}
