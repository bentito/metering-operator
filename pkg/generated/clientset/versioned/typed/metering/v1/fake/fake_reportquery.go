// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	meteringv1 "github.com/kubernetes-reporting/metering-operator/pkg/apis/metering/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeReportQueries implements ReportQueryInterface
type FakeReportQueries struct {
	Fake *FakeMeteringV1
	ns   string
}

var reportqueriesResource = schema.GroupVersionResource{Group: "metering.openshift.io", Version: "v1", Resource: "reportqueries"}

var reportqueriesKind = schema.GroupVersionKind{Group: "metering.openshift.io", Version: "v1", Kind: "ReportQuery"}

// Get takes name of the reportQuery, and returns the corresponding reportQuery object, and an error if there is any.
func (c *FakeReportQueries) Get(name string, options v1.GetOptions) (result *meteringv1.ReportQuery, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(reportqueriesResource, c.ns, name), &meteringv1.ReportQuery{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meteringv1.ReportQuery), err
}

// List takes label and field selectors, and returns the list of ReportQueries that match those selectors.
func (c *FakeReportQueries) List(opts v1.ListOptions) (result *meteringv1.ReportQueryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(reportqueriesResource, reportqueriesKind, c.ns, opts), &meteringv1.ReportQueryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &meteringv1.ReportQueryList{ListMeta: obj.(*meteringv1.ReportQueryList).ListMeta}
	for _, item := range obj.(*meteringv1.ReportQueryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested reportQueries.
func (c *FakeReportQueries) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(reportqueriesResource, c.ns, opts))

}

// Create takes the representation of a reportQuery and creates it.  Returns the server's representation of the reportQuery, and an error, if there is any.
func (c *FakeReportQueries) Create(reportQuery *meteringv1.ReportQuery) (result *meteringv1.ReportQuery, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(reportqueriesResource, c.ns, reportQuery), &meteringv1.ReportQuery{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meteringv1.ReportQuery), err
}

// Update takes the representation of a reportQuery and updates it. Returns the server's representation of the reportQuery, and an error, if there is any.
func (c *FakeReportQueries) Update(reportQuery *meteringv1.ReportQuery) (result *meteringv1.ReportQuery, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(reportqueriesResource, c.ns, reportQuery), &meteringv1.ReportQuery{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meteringv1.ReportQuery), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeReportQueries) UpdateStatus(reportQuery *meteringv1.ReportQuery) (*meteringv1.ReportQuery, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(reportqueriesResource, "status", c.ns, reportQuery), &meteringv1.ReportQuery{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meteringv1.ReportQuery), err
}

// Delete takes name of the reportQuery and deletes it. Returns an error if one occurs.
func (c *FakeReportQueries) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(reportqueriesResource, c.ns, name), &meteringv1.ReportQuery{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReportQueries) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(reportqueriesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &meteringv1.ReportQueryList{})
	return err
}

// Patch applies the patch and returns the patched reportQuery.
func (c *FakeReportQueries) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *meteringv1.ReportQuery, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(reportqueriesResource, c.ns, name, pt, data, subresources...), &meteringv1.ReportQuery{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meteringv1.ReportQuery), err
}
