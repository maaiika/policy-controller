// Copyright 2022 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
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

	v1alpha1 "github.com/sigstore/policy-controller/pkg/apis/policy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTrustRoots implements TrustRootInterface
type FakeTrustRoots struct {
	Fake *FakePolicyV1alpha1
}

var trustrootsResource = schema.GroupVersionResource{Group: "policy.sigstore.dev", Version: "v1alpha1", Resource: "trustroots"}

var trustrootsKind = schema.GroupVersionKind{Group: "policy.sigstore.dev", Version: "v1alpha1", Kind: "TrustRoot"}

// Get takes name of the trustRoot, and returns the corresponding trustRoot object, and an error if there is any.
func (c *FakeTrustRoots) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TrustRoot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(trustrootsResource, name), &v1alpha1.TrustRoot{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TrustRoot), err
}

// List takes label and field selectors, and returns the list of TrustRoots that match those selectors.
func (c *FakeTrustRoots) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TrustRootList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(trustrootsResource, trustrootsKind, opts), &v1alpha1.TrustRootList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TrustRootList{ListMeta: obj.(*v1alpha1.TrustRootList).ListMeta}
	for _, item := range obj.(*v1alpha1.TrustRootList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested trustRoots.
func (c *FakeTrustRoots) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(trustrootsResource, opts))
}

// Create takes the representation of a trustRoot and creates it.  Returns the server's representation of the trustRoot, and an error, if there is any.
func (c *FakeTrustRoots) Create(ctx context.Context, trustRoot *v1alpha1.TrustRoot, opts v1.CreateOptions) (result *v1alpha1.TrustRoot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(trustrootsResource, trustRoot), &v1alpha1.TrustRoot{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TrustRoot), err
}

// Update takes the representation of a trustRoot and updates it. Returns the server's representation of the trustRoot, and an error, if there is any.
func (c *FakeTrustRoots) Update(ctx context.Context, trustRoot *v1alpha1.TrustRoot, opts v1.UpdateOptions) (result *v1alpha1.TrustRoot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(trustrootsResource, trustRoot), &v1alpha1.TrustRoot{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TrustRoot), err
}

// Delete takes name of the trustRoot and deletes it. Returns an error if one occurs.
func (c *FakeTrustRoots) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(trustrootsResource, name, opts), &v1alpha1.TrustRoot{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTrustRoots) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(trustrootsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TrustRootList{})
	return err
}

// Patch applies the patch and returns the patched trustRoot.
func (c *FakeTrustRoots) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TrustRoot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(trustrootsResource, name, pt, data, subresources...), &v1alpha1.TrustRoot{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TrustRoot), err
}
