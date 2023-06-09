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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1beta1 "istio.io/client-go/pkg/apis/istio/v1beta1"
	istiov1beta1 "istio.io/client-go/pkg/applyconfiguration/istio/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTrafficLabels implements TrafficLabelInterface
type FakeTrafficLabels struct {
	Fake *FakeIstioV1beta1
	ns   string
}

var trafficlabelsResource = schema.GroupVersionResource{Group: "istio.alibabacloud.com", Version: "v1beta1", Resource: "trafficlabels"}

var trafficlabelsKind = schema.GroupVersionKind{Group: "istio.alibabacloud.com", Version: "v1beta1", Kind: "TrafficLabel"}

// Get takes name of the trafficLabel, and returns the corresponding trafficLabel object, and an error if there is any.
func (c *FakeTrafficLabels) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.TrafficLabel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(trafficlabelsResource, c.ns, name), &v1beta1.TrafficLabel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.TrafficLabel), err
}

// List takes label and field selectors, and returns the list of TrafficLabels that match those selectors.
func (c *FakeTrafficLabels) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.TrafficLabelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(trafficlabelsResource, trafficlabelsKind, c.ns, opts), &v1beta1.TrafficLabelList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.TrafficLabelList{ListMeta: obj.(*v1beta1.TrafficLabelList).ListMeta}
	for _, item := range obj.(*v1beta1.TrafficLabelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested trafficLabels.
func (c *FakeTrafficLabels) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(trafficlabelsResource, c.ns, opts))

}

// Create takes the representation of a trafficLabel and creates it.  Returns the server's representation of the trafficLabel, and an error, if there is any.
func (c *FakeTrafficLabels) Create(ctx context.Context, trafficLabel *v1beta1.TrafficLabel, opts v1.CreateOptions) (result *v1beta1.TrafficLabel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(trafficlabelsResource, c.ns, trafficLabel), &v1beta1.TrafficLabel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.TrafficLabel), err
}

// Update takes the representation of a trafficLabel and updates it. Returns the server's representation of the trafficLabel, and an error, if there is any.
func (c *FakeTrafficLabels) Update(ctx context.Context, trafficLabel *v1beta1.TrafficLabel, opts v1.UpdateOptions) (result *v1beta1.TrafficLabel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(trafficlabelsResource, c.ns, trafficLabel), &v1beta1.TrafficLabel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.TrafficLabel), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTrafficLabels) UpdateStatus(ctx context.Context, trafficLabel *v1beta1.TrafficLabel, opts v1.UpdateOptions) (*v1beta1.TrafficLabel, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(trafficlabelsResource, "status", c.ns, trafficLabel), &v1beta1.TrafficLabel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.TrafficLabel), err
}

// Delete takes name of the trafficLabel and deletes it. Returns an error if one occurs.
func (c *FakeTrafficLabels) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(trafficlabelsResource, c.ns, name, opts), &v1beta1.TrafficLabel{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTrafficLabels) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(trafficlabelsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.TrafficLabelList{})
	return err
}

// Patch applies the patch and returns the patched trafficLabel.
func (c *FakeTrafficLabels) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.TrafficLabel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(trafficlabelsResource, c.ns, name, pt, data, subresources...), &v1beta1.TrafficLabel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.TrafficLabel), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied trafficLabel.
func (c *FakeTrafficLabels) Apply(ctx context.Context, trafficLabel *istiov1beta1.TrafficLabelApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.TrafficLabel, err error) {
	if trafficLabel == nil {
		return nil, fmt.Errorf("trafficLabel provided to Apply must not be nil")
	}
	data, err := json.Marshal(trafficLabel)
	if err != nil {
		return nil, err
	}
	name := trafficLabel.Name
	if name == nil {
		return nil, fmt.Errorf("trafficLabel.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(trafficlabelsResource, c.ns, *name, types.ApplyPatchType, data), &v1beta1.TrafficLabel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.TrafficLabel), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeTrafficLabels) ApplyStatus(ctx context.Context, trafficLabel *istiov1beta1.TrafficLabelApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.TrafficLabel, err error) {
	if trafficLabel == nil {
		return nil, fmt.Errorf("trafficLabel provided to Apply must not be nil")
	}
	data, err := json.Marshal(trafficLabel)
	if err != nil {
		return nil, err
	}
	name := trafficLabel.Name
	if name == nil {
		return nil, fmt.Errorf("trafficLabel.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(trafficlabelsResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1beta1.TrafficLabel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.TrafficLabel), err
}
