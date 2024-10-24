package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetCustomPolicyForPermissionSetResponse Response Object
type GetCustomPolicyForPermissionSetResponse struct {

	// 附加到权限集的自定义身份策略
	CustomPolicy   *string `json:"custom_policy,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GetCustomPolicyForPermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetCustomPolicyForPermissionSetResponse struct{}"
	}

	return strings.Join([]string{"GetCustomPolicyForPermissionSetResponse", string(data)}, " ")
}
