package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UntagDeviceRequest struct {
	// **参数说明**：实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。

	InstanceId *string `json:"Instance-Id,omitempty"`

	Body *UnbindTagsDto `json:"body,omitempty"`
}

func (o UntagDeviceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UntagDeviceRequest struct{}"
	}

	return strings.Join([]string{"UntagDeviceRequest", string(data)}, " ")
}
