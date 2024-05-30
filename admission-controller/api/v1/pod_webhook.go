/*
Copyright 2024.

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

package v1

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// // log is for logging in this package.
// var podlog = logf.Log.WithName("pod-resource")

// // SetupWebhookWithManager will setup the manager to manage the webhooks
// func (r *Pod) SetupWebhookWithManager(mgr ctrl.Manager) error {
// 	return ctrl.NewWebhookManagedBy(mgr).
// 		For(r).
// 		Complete()
// }

// // TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// // +kubebuilder:webhook:path=/mutate-v1-pod,mutating=true,failurePolicy=fail,groups="",resources=pods,verbs=create;update,versions=v1,name=mpod.kb.io
// var _ webhook.Defaulter = &Pod{}

// // Default implements webhook.Defaulter so a webhook will be registered for the type
// func (r *Pod) Default() {
// 	podlog.Info("default", "name", r.Name)

// 	// TODO(user): fill in your defaulting logic.
// }

// // TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// // NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// // Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// // +kubebuilder:webhook:path=/validate-core-v1-pod,mutating=false,failurePolicy=fail,sideEffects=None,groups=core,resources=pods,verbs=create;update,versions=v1,name=vpod.kb.io,admissionReviewVersions=v1

// var _ webhook.Validator = &Pod{}

// // ValidateCreate implements webhook.Validator so a webhook will be registered for the type
// func (r *Pod) ValidateCreate() (admission.Warnings, error) {
// 	podlog.Info("validate create", "name", r.Name)

// 	// TODO(user): fill in your validation logic upon object creation.
// 	return nil, nil
// }

// // ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
// func (r *Pod) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
// 	podlog.Info("validate update", "name", r.Name)

// 	// TODO(user): fill in your validation logic upon object update.
// 	return nil, nil
// }

// // ValidateDelete implements webhook.Validator so a webhook will be registered for the type
// func (r *Pod) ValidateDelete() (admission.Warnings, error) {
// 	podlog.Info("validate delete", "name", r.Name)

// 	// TODO(user): fill in your validation logic upon object deletion.
// 	return nil, nil
// }

type PodAnnotator struct {
	Client  client.Client
	Decoder admission.Decoder
}

func (a *PodAnnotator) Handle(ctx context.Context, req admission.Request) admission.Response {
	pod := &corev1.Pod{}
	err := a.Decoder.Decode(req, pod)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	// mutate the fields in pod
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	logger = logger.With(zap.Any("req", req))
	logger.Info("printing admission controller req")

	marshaledPod, err := json.Marshal(pod)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError, err)
	}
	return admission.PatchResponseFromRaw(req.Object.Raw, marshaledPod)
}
