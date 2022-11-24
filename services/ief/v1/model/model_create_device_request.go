package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDeviceRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	Body *EdgemgrDevices `json:"body,omitempty"`
}

func (o CreateDeviceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeviceRequest struct{}"
	}

	return strings.Join([]string{"CreateDeviceRequest", string(data)}, " ")
}
