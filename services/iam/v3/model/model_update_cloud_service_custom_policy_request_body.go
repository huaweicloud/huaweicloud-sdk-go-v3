package model

import (
	"encoding/json"

	"strings"
)

//
type UpdateCloudServiceCustomPolicyRequestBody struct {
	Role *ServicePolicyRoleOption `json:"role"`
}

func (o UpdateCloudServiceCustomPolicyRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateCloudServiceCustomPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateCloudServiceCustomPolicyRequestBody", string(data)}, " ")
}
