/*
Copyright The Flagger Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta2

import (
	"context"
	time "time"

	appmeshv1beta2 "github.com/fluxcd/flagger/pkg/apis/appmesh/v1beta2"
	versioned "github.com/fluxcd/flagger/pkg/client/clientset/versioned"
	internalinterfaces "github.com/fluxcd/flagger/pkg/client/informers/externalversions/internalinterfaces"
	v1beta2 "github.com/fluxcd/flagger/pkg/client/listers/appmesh/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VirtualRouterInformer provides access to a shared informer and lister for
// VirtualRouters.
type VirtualRouterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta2.VirtualRouterLister
}

type virtualRouterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVirtualRouterInformer constructs a new informer for VirtualRouter type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVirtualRouterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVirtualRouterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVirtualRouterInformer constructs a new informer for VirtualRouter type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVirtualRouterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppmeshV1beta2().VirtualRouters(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppmeshV1beta2().VirtualRouters(namespace).Watch(context.TODO(), options)
			},
		},
		&appmeshv1beta2.VirtualRouter{},
		resyncPeriod,
		indexers,
	)
}

func (f *virtualRouterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVirtualRouterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *virtualRouterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&appmeshv1beta2.VirtualRouter{}, f.defaultInformer)
}

func (f *virtualRouterInformer) Lister() v1beta2.VirtualRouterLister {
	return v1beta2.NewVirtualRouterLister(f.Informer().GetIndexer())
}
