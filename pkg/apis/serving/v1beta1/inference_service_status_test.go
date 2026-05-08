package v1beta1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransitionStatusConstants(t *testing.T) {
	assert.Equal(t, TransitionStatus("UpToDate"), UpToDate)
	assert.Equal(t, TransitionStatus("InProgress"), InProgress)
	assert.Equal(t, TransitionStatus("BlockedByFailure"), BlockedByFailure)
	assert.Equal(t, TransitionStatus("InvalidSpec"), InvalidSpec)
}

func TestModelStateConstants(t *testing.T) {
	assert.Equal(t, ModelState("Pending"), Pending)
	assert.Equal(t, ModelState("Standby"), Standby)
	assert.Equal(t, ModelState("Loading"), Loading)
	assert.Equal(t, ModelState("Loaded"), Loaded)
	assert.Equal(t, ModelState("FailedToLoad"), FailedToLoad)
}

func TestFailureReasonConstants(t *testing.T) {
	assert.Equal(t, FailureReason("ModelLoadFailed"), ModelLoadFailed)
	assert.Equal(t, FailureReason("RuntimeUnhealthy"), RuntimeUnhealthy)
	assert.Equal(t, FailureReason("NoSupportingRuntime"), NoSupportingRuntime)
	assert.Equal(t, FailureReason("RuntimeNotRecognized"), RuntimeNotRecognized)
	assert.Equal(t, FailureReason("InvalidPredictorSpec"), InvalidPredictorSpec)
}

func TestModelStatusDefaults(t *testing.T) {
	// Verify that a zero-value ModelStatus has no transition status, revision states,
	// or failure info set — important for distinguishing freshly created resources.
	status := ModelStatus{}
	assert.Empty(t, status.TransitionStatus)
	assert.Nil(t, status.ModelRevisionStates)
	assert.Nil(t, status.LastFailureInfo)
}

func TestFailureInfoFields(t *testing.T) {
	msg := "failed to load model"
	revName := "rev-001"
	var ts int64 = 1700000000
	info := FailureInfo{
		Reason:            ModelLoadFailed,
		Message:           msg,
		ModelRevisionName: revName,
		Time:              &ts,
	}
	assert.Equal(t, ModelLoadFailed, info.Reason)
	assert.Equal(t, msg, info.Message)
	assert.Equal(t, revName, info.ModelRevisionName)
	assert.Equal(t, int64(1700000000), *info.Time)

	// Also verify that a nil Time pointer is handled gracefully (no panic on dereference).
	infoNoTime := FailureInfo{
		Reason:  RuntimeUnhealthy,
		Message: "runtime pod crashed",
	}
	assert.Nil(t, infoNoTime.Time)
}

func TestModelRevisionStates(t *testing.T) {
	states := ModelRevisionStates{
		ActiveModelState: Loaded,
		TargetModelState: Loading,
	}
	assert.Equal(t, Loaded, states.ActiveModelState)
	assert.Equal(t, Loading, states.TargetModelState)
}

// TestModelRevisionStatesOnlyActive checks the case where only the active model
// state is set and the target state is empty (i.e., no rollout in progress).
func TestModelRevisionStatesOnlyActive(t *testing.T) {
	states := ModelRevisionStates{
		ActiveModelState: Loaded,
	}
	assert.Equal(t, Loaded, states.ActiveModelState)
	assert.Empty(t, states.TargetModelState)
}

func TestInferenceServiceStatusComponents(t *testing.T) {
	status := InferenceServiceStatus{
		Components: map[ComponentType]ComponentStatusSpec{},
	}
	assert.NotNil(t, status.Components)
	assert.Len(t, status.Components, 0)
}

func TestInferenceServiceConditionTypes(t *testing.T) {
	assert.Equal(t, InferenceServiceConditionType("PredictorReady"), PredictorReady)
	assert.Equal(t, InferenceServiceConditionType("TransformerReady"), TransformerReady)
	assert.Equal(t, InferenceServiceConditionType("ExplainerReady"), ExplainerReady)
	assert.Equal(t, InferenceServiceConditionType("IngressReady"), IngressReady)
	assert.Equal(t, InferenceServiceConditionType("Ready"), InferenceServiceReady)
}
