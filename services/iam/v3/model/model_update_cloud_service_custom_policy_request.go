package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateCloudServiceCustomPolicyRequest struct {
	RoleId string `json:"role_id"`

	Body *UpdateCloudServiceCustomPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateCloudServiceCustomPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateCloudServiceCustomPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateCloudServiceCustomPolicyRequest", string(data)}, " ")
}
