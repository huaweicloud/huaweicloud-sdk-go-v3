package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddDeviceRequest struct {
	// **参数说明**：实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。

	InstanceId *string `json:"Instance-Id,omitempty"`

	Body *AddDevice `json:"body,omitempty"`
}

func (o AddDeviceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDeviceRequest struct{}"
	}

	return strings.Join([]string{"AddDeviceRequest", string(data)}, " ")
}
