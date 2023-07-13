package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDeviceTemplatesResponse Response Object
type ListDeviceTemplatesResponse struct {

	// 终端设备属性
	DeviceTemplates *[]EdgemgrDevice `json:"device_templates,omitempty"`

	// 模板数量
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDeviceTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeviceTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListDeviceTemplatesResponse", string(data)}, " ")
}
