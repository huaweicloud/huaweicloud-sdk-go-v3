package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetResourcesRenewConfigResponse Response Object
type SetResourcesRenewConfigResponse struct {

	// |参数名称：返回码| |参数的约束及描述：该参数必填，且只允许字符串|
	ErrorCode *string `json:"error_code,omitempty"`

	// |参数名称：返回码描述| |参数的约束及描述：该参数必填，且只允许字符串|
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetResourcesRenewConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetResourcesRenewConfigResponse struct{}"
	}

	return strings.Join([]string{"SetResourcesRenewConfigResponse", string(data)}, " ")
}
