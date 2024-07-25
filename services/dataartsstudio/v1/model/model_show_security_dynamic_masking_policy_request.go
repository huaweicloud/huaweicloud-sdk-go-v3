package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityDynamicMaskingPolicyRequest Request Object
type ShowSecurityDynamicMaskingPolicyRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 动态脱敏策略id。
	Id string `json:"id"`
}

func (o ShowSecurityDynamicMaskingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityDynamicMaskingPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowSecurityDynamicMaskingPolicyRequest", string(data)}, " ")
}
