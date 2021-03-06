// Code generated by k8s code-generator DO NOT EDIT.

/*
Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package mocks

import (
	clientset "github.com/GoogleCloudPlatform/spark-on-k8s-operator/pkg/client/clientset/versioned"
	sparkoperatorv1beta1 "github.com/GoogleCloudPlatform/spark-on-k8s-operator/pkg/client/clientset/versioned/typed/sparkoperator.k8s.io/v1beta1"
	sparkoperatorv1beta2 "github.com/GoogleCloudPlatform/spark-on-k8s-operator/pkg/client/clientset/versioned/typed/sparkoperator.k8s.io/v1beta2"
	fakesparkoperatorv1beta1 "github.com/gojek/merlin/batch/mocks/v1beta1/mocks"
	fakesparkoperatorv1beta2 "github.com/gojek/merlin/batch/mocks/v1beta2/mocks"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

var _ clientset.Interface = &Clientset{}

// SparkoperatorV1beta1 retrieves the SparkoperatorV1beta1Client
func (c *Clientset) SparkoperatorV1beta1() sparkoperatorv1beta1.SparkoperatorV1beta1Interface {
	return &fakesparkoperatorv1beta1.FakeSparkoperatorV1beta1{Fake: &c.Fake}
}

// SparkoperatorV1beta2 retrieves the SparkoperatorV1beta2Client
func (c *Clientset) SparkoperatorV1beta2() sparkoperatorv1beta2.SparkoperatorV1beta2Interface {
	return &fakesparkoperatorv1beta2.FakeSparkoperatorV1beta2{Fake: &c.Fake}
}

// Sparkoperator retrieves the SparkoperatorV1beta2Client
func (c *Clientset) Sparkoperator() sparkoperatorv1beta2.SparkoperatorV1beta2Interface {
	return &fakesparkoperatorv1beta2.FakeSparkoperatorV1beta2{Fake: &c.Fake}
}
