package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type UpdateCloudServiceCustomPolicyRequestBody struct {
	Role *ServicePolicyRoleOption `json:"role"`
}

func (o UpdateCloudServiceCustomPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCloudServiceCustomPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateCloudServiceCustomPolicyRequestBody", string(data)}, " ")
}
