package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FailResourceInfo struct {

	// |参数名称：错误码| |参数约束及描述：错误码|
	ErrorCode *string `json:"error_code,omitempty"`

	// |参数名称：错误描述| |参数约束及描述：错误描述|
	ErrorMsg *string `json:"error_msg,omitempty"`

	// |参数名称：资源ID| |参数约束及描述：资源ID|
	ResourceId *string `json:"resource_id,omitempty"`
}

func (o FailResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailResourceInfo struct{}"
	}

	return strings.Join([]string{"FailResourceInfo", string(data)}, " ")
}
