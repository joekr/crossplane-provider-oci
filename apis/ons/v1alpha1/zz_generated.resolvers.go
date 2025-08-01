/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/oracle/provider-oci/apis/identity/v1alpha1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this NotificationTopic.
func (mg *NotificationTopic) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CompartmentIDRef,
		Selector:     mg.Spec.ForProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &v1alpha1.CompartmentList{},
			Managed: &v1alpha1.Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CompartmentID")
	}
	mg.Spec.ForProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CompartmentIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.CompartmentIDRef,
		Selector:     mg.Spec.InitProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &v1alpha1.CompartmentList{},
			Managed: &v1alpha1.Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CompartmentID")
	}
	mg.Spec.InitProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CompartmentIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Subscription.
func (mg *Subscription) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CompartmentIDRef,
		Selector:     mg.Spec.ForProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &v1alpha1.CompartmentList{},
			Managed: &v1alpha1.Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CompartmentID")
	}
	mg.Spec.ForProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CompartmentIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TopicID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TopicIDRef,
		Selector:     mg.Spec.ForProvider.TopicIDSelector,
		To: reference.To{
			List:    &NotificationTopicList{},
			Managed: &NotificationTopic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TopicID")
	}
	mg.Spec.ForProvider.TopicID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TopicIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.CompartmentIDRef,
		Selector:     mg.Spec.InitProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &v1alpha1.CompartmentList{},
			Managed: &v1alpha1.Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CompartmentID")
	}
	mg.Spec.InitProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CompartmentIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TopicID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.TopicIDRef,
		Selector:     mg.Spec.InitProvider.TopicIDSelector,
		To: reference.To{
			List:    &NotificationTopicList{},
			Managed: &NotificationTopic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.TopicID")
	}
	mg.Spec.InitProvider.TopicID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TopicIDRef = rsp.ResolvedReference

	return nil
}
