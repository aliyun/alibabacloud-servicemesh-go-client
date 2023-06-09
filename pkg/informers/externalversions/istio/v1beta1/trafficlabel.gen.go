// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	istiov1beta1 "istio.io/client-go/pkg/apis/istio/v1beta1"
	versioned "istio.io/client-go/pkg/clientset/versioned"
	internalinterfaces "istio.io/client-go/pkg/informers/externalversions/internalinterfaces"
	v1beta1 "istio.io/client-go/pkg/listers/istio/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TrafficLabelInformer provides access to a shared informer and lister for
// TrafficLabels.
type TrafficLabelInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.TrafficLabelLister
}

type trafficLabelInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTrafficLabelInformer constructs a new informer for TrafficLabel type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTrafficLabelInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTrafficLabelInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTrafficLabelInformer constructs a new informer for TrafficLabel type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTrafficLabelInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IstioV1beta1().TrafficLabels(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IstioV1beta1().TrafficLabels(namespace).Watch(context.TODO(), options)
			},
		},
		&istiov1beta1.TrafficLabel{},
		resyncPeriod,
		indexers,
	)
}

func (f *trafficLabelInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTrafficLabelInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *trafficLabelInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&istiov1beta1.TrafficLabel{}, f.defaultInformer)
}

func (f *trafficLabelInformer) Lister() v1beta1.TrafficLabelLister {
	return v1beta1.NewTrafficLabelLister(f.Informer().GetIndexer())
}
