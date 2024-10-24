package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PutCustomPolicyToPermissionSetReqBody the request body of PutCustomPolicyToPermissionSet
type PutCustomPolicyToPermissionSetReqBody struct {

	// 要附加到权限集的自定义身份策略
	CustomPolicy string `json:"custom_policy"`
}

func (o PutCustomPolicyToPermissionSetReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutCustomPolicyToPermissionSetReqBody struct{}"
	}

	return strings.Join([]string{"PutCustomPolicyToPermissionSetReqBody", string(data)}, " ")
}
