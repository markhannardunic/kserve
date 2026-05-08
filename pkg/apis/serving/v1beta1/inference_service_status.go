package v1beta1

import (
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

// InferenceServiceConditionType defines the condition types for InferenceService
type InferenceServiceConditionType string

const (
	// PredictorReady defines the condition for predictor readiness
	PredictorReady InferenceServiceConditionType = "PredictorReady"
	// TransformerReady defines the condition for transformer readiness
	TransformerReady InferenceServiceConditionType = "TransformerReady"
	// ExplainerReady defines the condition for explainer readiness
	ExplainerReady InferenceServiceConditionType = "ExplainerReady"
	// IngressReady defines the condition for ingress readiness
	IngressReady InferenceServiceConditionType = "IngressReady"
	// LatestDeploymentReady defines the condition for latest deployment readiness
	LatestDeploymentReady InferenceServiceConditionType = "LatestDeploymentReady"
)

// InferenceServiceStatus defines the observed state of InferenceService
type InferenceServiceStatus struct {
	duckv1.Status `json:",inline"`
	// URL is the endpoint url of the inference service
	// +optional
	URL *apis.URL `json:"url,omitempty"`
	// ModelStatus holds the per-model status
	// +optional
	ModelStatus ModelStatus `json:"modelStatus,omitempty"`
	// Components holds the per-component status
	// +optional
	Components map[ComponentType]ComponentStatusSpec `json:"components,omitempty"`
}

// ModelStatus describes the status of a model
type ModelStatus struct {
	// TransitionStatus is the status of the model transition
	TransitionStatus TransitionStatus `json:"transitionStatus,omitempty"`
	// ModelRevisionStates holds the states of model revisions
	ModelRevisionStates *ModelRevisionStates `json:"modelRevisionStates,omitempty"`
	// LastFailureInfo holds info about the last failure
	LastFailureInfo *FailureInfo `json:"lastFailureInfo,omitempty"`
}

// TransitionStatus enum
type TransitionStatus string

const (
	UpToDate         TransitionStatus = "UpToDate"
	InProgress       TransitionStatus = "InProgress"
	BlockedByFailure TransitionStatus = "BlockedByFailure"
	InvalidSpec      TransitionStatus = "InvalidSpec"
)

// ModelRevisionStates holds states for active and target model revisions
type ModelRevisionStates struct {
	ActiveModelState ModelState `json:"activeModelState,omitempty"`
	TargetModelState ModelState `json:"targetModelState,omitempty"`
}

// ModelState enum
type ModelState string

const (
	Pending ModelState = "Pending"
	Standby ModelState = "Standby"
	Loading ModelState = "Loading"
	Loaded  ModelState = "Loaded"
	// FailedToLoad indicates the model could not be loaded; check LastFailureInfo for details
	FailedToLoad ModelState = "FailedToLoad"
)

// FailureReason is the reason for a model failure
type FailureReason string

const (
	// ModelLoadFailed indicates the model failed to load (e.g. bad weights or OOM)
	ModelLoadFailed FailureReason = "ModelLoadFailed"
	// RuntimeUnhealthy indicates the model runtime pod is not healthy
	RuntimeUnhealthy FailureReason = "RuntimeUnhealthy"
	// NoSupportingRuntime indicates no runtime supports the requested model type
	NoSupportingRuntime FailureReason = "NoSupportingRuntime"
	// InvalidPredictorSpec indicates the predictor spec is invalid
	InvalidPredictorSpec FailureReason = "InvalidPredictorSpec"
)

// FailureInfo holds information about a failure
type FailureInfo struct {
	Reason            FailureReason `json:"reason,omitempty"`
	Message           string        `json:"message,omitempty"`
	ModelRevisionName string        `json:"modelRevisionName,omitempty"`
	// Time is the Unix timestamp (seconds) when the failure occurred
	Time *int64 `json:"time,omitempty"`
}
