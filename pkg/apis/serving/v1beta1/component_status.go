package v1beta1

import "knative.dev/pkg/apis"

// ComponentType defines the types of components in an InferenceService
type ComponentType string

const (
	// PredictorComponent is the predictor component
	PredictorComponent ComponentType = "predictor"
	// TransformerComponent is the transformer component
	TransformerComponent ComponentType = "transformer"
	// ExplainerComponent is the explainer component
	ExplainerComponent ComponentType = "explainer"
)

// ComponentStatusSpec describes the status of a single component
type ComponentStatusSpec struct {
	// LatestReadyRevision is the name of the latest ready revision
	// +optional
	LatestReadyRevision string `json:"latestReadyRevision,omitempty"`
	// LatestCreatedRevision is the name of the latest created revision
	// +optional
	LatestCreatedRevision string `json:"latestCreatedRevision,omitempty"`
	// PreviousRolledoutRevision is the name of the previous rolledout revision
	// +optional
	PreviousRolledoutRevision string `json:"previousRolledoutRevision,omitempty"`
	// LatestRolledoutRevision is the name of the latest rolledout revision
	// +optional
	LatestRolledoutRevision string `json:"latestRolledoutRevision,omitempty"`
	// Traffic holds the traffic split between revisions
	// +optional
	Traffic []RevisionTrafficStatus `json:"traffic,omitempty"`
	// URL is the component endpoint url
	// +optional
	URL *apis.URL `json:"url,omitempty"`
	// RestURL is the REST endpoint url
	// +optional
	RestURL *apis.URL `json:"restUrl,omitempty"`
	// GrpcURL is the gRPC endpoint url
	// +optional
	GrpcURL *apis.URL `json:"grpcUrl,omitempty"`
}

// RevisionTrafficStatus describes traffic routed to a revision
type RevisionTrafficStatus struct {
	// RevisionName is the name of the revision
	RevisionName string `json:"revisionName,omitempty"`
	// Percent is the percentage of traffic routed to the revision.
	// Valid range is 0 to 100 (inclusive); all revision percents must sum to 100.
	// +optional
	Percent *int64 `json:"percent,omitempty"`
	// Tag is the name tag associated with the revision
	// +optional
	Tag string `json:"tag,omitempty"`
	// URL is the URL for the revision traffic
	// +optional
	URL *apis.URL `json:"url,omitempty"`
	// LatestRevision indicates if this is the latest revision
	// +optional
	LatestRevision *bool `json:"latestRevision,omitempty"`
}
