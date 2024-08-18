/*
Copyright 2022.

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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/odigos-io/odigos/api/actions/v1alpha1"
	actionsv1alpha1 "github.com/odigos-io/odigos/api/generated/actions/applyconfiguration/actions/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAddClusterInfos implements AddClusterInfoInterface
type FakeAddClusterInfos struct {
	Fake *FakeActionsV1alpha1
	ns   string
}

var addclusterinfosResource = v1alpha1.SchemeGroupVersion.WithResource("addclusterinfos")

var addclusterinfosKind = v1alpha1.SchemeGroupVersion.WithKind("AddClusterInfo")

// Get takes name of the addClusterInfo, and returns the corresponding addClusterInfo object, and an error if there is any.
func (c *FakeAddClusterInfos) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AddClusterInfo, err error) {
	emptyResult := &v1alpha1.AddClusterInfo{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(addclusterinfosResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.AddClusterInfo), err
}

// List takes label and field selectors, and returns the list of AddClusterInfos that match those selectors.
func (c *FakeAddClusterInfos) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AddClusterInfoList, err error) {
	emptyResult := &v1alpha1.AddClusterInfoList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(addclusterinfosResource, addclusterinfosKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AddClusterInfoList{ListMeta: obj.(*v1alpha1.AddClusterInfoList).ListMeta}
	for _, item := range obj.(*v1alpha1.AddClusterInfoList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested addClusterInfos.
func (c *FakeAddClusterInfos) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(addclusterinfosResource, c.ns, opts))

}

// Create takes the representation of a addClusterInfo and creates it.  Returns the server's representation of the addClusterInfo, and an error, if there is any.
func (c *FakeAddClusterInfos) Create(ctx context.Context, addClusterInfo *v1alpha1.AddClusterInfo, opts v1.CreateOptions) (result *v1alpha1.AddClusterInfo, err error) {
	emptyResult := &v1alpha1.AddClusterInfo{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(addclusterinfosResource, c.ns, addClusterInfo, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.AddClusterInfo), err
}

// Update takes the representation of a addClusterInfo and updates it. Returns the server's representation of the addClusterInfo, and an error, if there is any.
func (c *FakeAddClusterInfos) Update(ctx context.Context, addClusterInfo *v1alpha1.AddClusterInfo, opts v1.UpdateOptions) (result *v1alpha1.AddClusterInfo, err error) {
	emptyResult := &v1alpha1.AddClusterInfo{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(addclusterinfosResource, c.ns, addClusterInfo, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.AddClusterInfo), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAddClusterInfos) UpdateStatus(ctx context.Context, addClusterInfo *v1alpha1.AddClusterInfo, opts v1.UpdateOptions) (result *v1alpha1.AddClusterInfo, err error) {
	emptyResult := &v1alpha1.AddClusterInfo{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(addclusterinfosResource, "status", c.ns, addClusterInfo, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.AddClusterInfo), err
}

// Delete takes name of the addClusterInfo and deletes it. Returns an error if one occurs.
func (c *FakeAddClusterInfos) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(addclusterinfosResource, c.ns, name, opts), &v1alpha1.AddClusterInfo{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAddClusterInfos) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(addclusterinfosResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AddClusterInfoList{})
	return err
}

// Patch applies the patch and returns the patched addClusterInfo.
func (c *FakeAddClusterInfos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AddClusterInfo, err error) {
	emptyResult := &v1alpha1.AddClusterInfo{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(addclusterinfosResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.AddClusterInfo), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied addClusterInfo.
func (c *FakeAddClusterInfos) Apply(ctx context.Context, addClusterInfo *actionsv1alpha1.AddClusterInfoApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.AddClusterInfo, err error) {
	if addClusterInfo == nil {
		return nil, fmt.Errorf("addClusterInfo provided to Apply must not be nil")
	}
	data, err := json.Marshal(addClusterInfo)
	if err != nil {
		return nil, err
	}
	name := addClusterInfo.Name
	if name == nil {
		return nil, fmt.Errorf("addClusterInfo.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.AddClusterInfo{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(addclusterinfosResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.AddClusterInfo), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeAddClusterInfos) ApplyStatus(ctx context.Context, addClusterInfo *actionsv1alpha1.AddClusterInfoApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.AddClusterInfo, err error) {
	if addClusterInfo == nil {
		return nil, fmt.Errorf("addClusterInfo provided to Apply must not be nil")
	}
	data, err := json.Marshal(addClusterInfo)
	if err != nil {
		return nil, err
	}
	name := addClusterInfo.Name
	if name == nil {
		return nil, fmt.Errorf("addClusterInfo.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.AddClusterInfo{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(addclusterinfosResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.AddClusterInfo), err
}
