package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteDeviceControlsReleaseResponse Response Object
type ExecuteDeviceControlsReleaseResponse struct {

	// 属性设置的响应码，具体为实际设备返回的响应码
	ResultCode *int32 `json:"result_code,omitempty"`

	// 属性设置的描述，具体为实际设备返回的描述
	ResultDesc     *string `json:"result_desc,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteDeviceControlsReleaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDeviceControlsReleaseResponse struct{}"
	}

	return strings.Join([]string{"ExecuteDeviceControlsReleaseResponse", string(data)}, " ")
}
