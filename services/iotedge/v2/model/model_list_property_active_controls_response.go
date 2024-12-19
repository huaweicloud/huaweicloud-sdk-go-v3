package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPropertyActiveControlsResponse Response Object
type ListPropertyActiveControlsResponse struct {

	// 设备id
	DeviceId *string `json:"device_id,omitempty"`

	// 设备的服务id
	ServiceId *string `json:"service_id,omitempty"`

	// 设备的属性
	Property *string `json:"property,omitempty"`

	// 正在执行中的控制列表
	ActiveControls *[]ActiveControlRspDto `json:"active_controls,omitempty"`

	// 结果响应码，大于等于400表示失败，小于400表示成功
	ResultCode *int32 `json:"result_code,omitempty"`

	// 结果描述
	ResultDesc     *string `json:"result_desc,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPropertyActiveControlsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPropertyActiveControlsResponse struct{}"
	}

	return strings.Join([]string{"ListPropertyActiveControlsResponse", string(data)}, " ")
}
