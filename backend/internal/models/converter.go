package models

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// ConvertToMetadata extracts metadata from an unstructured object
func ConvertToMetadata(obj *unstructured.Unstructured) Metadata {
	return Metadata{
		Name:              obj.GetName(),
		Namespace:         obj.GetNamespace(),
		UID:               string(obj.GetUID()),
		Labels:            obj.GetLabels(),
		Annotations:       obj.GetAnnotations(),
		CreationTimestamp: obj.GetCreationTimestamp().Time,
	}
}

// ConvertToBaseResource creates a BaseResource from an unstructured object
func ConvertToBaseResource(obj *unstructured.Unstructured, scope ResourceScope) BaseResource {
	return BaseResource{
		Kind:       obj.GetKind(),
		APIVersion: obj.GetAPIVersion(),
		Metadata:   ConvertToMetadata(obj),
		Scope:      scope,
	}
}

// ConvertConditions extracts conditions from status
func ConvertConditions(obj map[string]interface{}) []Condition {
	conditionsRaw, found := obj["conditions"]
	if !found {
		return []Condition{}
	}

	conditionsList, ok := conditionsRaw.([]interface{})
	if !ok {
		return []Condition{}
	}

	var conditions []Condition
	for _, c := range conditionsList {
		condMap, ok := c.(map[string]interface{})
		if !ok {
			continue
		}

		condition := Condition{
			Type:   getStringField(condMap, "type"),
			Status: getStringField(condMap, "status"),
			Reason: getStringField(condMap, "reason"),
			Message: getStringField(condMap, "message"),
		}

		// Parse lastTransitionTime
		if timeStr, ok := condMap["lastTransitionTime"].(string); ok {
			if t, err := time.Parse(time.RFC3339, timeStr); err == nil {
				condition.LastTransitionTime = t
			}
		}

		conditions = append(conditions, condition)
	}

	return conditions
}

// IsResourceReady checks if a resource is ready based on its conditions
func IsResourceReady(conditions []Condition) bool {
	for _, cond := range conditions {
		if cond.Type == "Ready" && cond.Status == "True" {
			return true
		}
	}
	return false
}

// ConvertToResourceStatus extracts status from an unstructured object
func ConvertToResourceStatus(obj *unstructured.Unstructured) ResourceStatus {
	status, found, _ := unstructured.NestedMap(obj.Object, "status")
	if !found {
		return ResourceStatus{
			Conditions: []Condition{},
			Ready:      false,
		}
	}

	conditions := ConvertConditions(status)
	return ResourceStatus{
		Conditions: conditions,
		Ready:      IsResourceReady(conditions),
	}
}

// Helper function to safely get string field
func getStringField(obj map[string]interface{}, field string) string {
	if val, ok := obj[field].(string); ok {
		return val
	}
	return ""
}

// Helper function to safely get int field
func getIntField(obj map[string]interface{}, field string) int {
	if val, ok := obj[field].(int); ok {
		return val
	}
	if val, ok := obj[field].(int64); ok {
		return int(val)
	}
	if val, ok := obj[field].(float64); ok {
		return int(val)
	}
	return 0
}

// ConvertTimeString converts a string timestamp to time.Time
func ConvertTimeString(timeStr string) time.Time {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		// Try metav1.Time format
		mt := &metav1.Time{}
		if err := mt.UnmarshalQueryParameter(timeStr); err == nil {
			return mt.Time
		}
		return time.Time{}
	}
	return t
}
