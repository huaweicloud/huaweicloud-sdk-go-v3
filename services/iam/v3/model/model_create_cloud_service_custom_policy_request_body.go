package model

import (
	"encoding/json"

	"strings"
)

//
type CreateCloudServiceCustomPolicyRequestBody struct {
	Role *ServicePolicyRoleOption `json:"role"`
}

func (o CreateCloudServiceCustomPolicyRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateCloudServiceCustomPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCloudServiceCustomPolicyRequestBody", string(data)}, " ")
}
