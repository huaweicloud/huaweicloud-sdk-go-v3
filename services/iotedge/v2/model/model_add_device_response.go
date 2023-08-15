package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDeviceResponse Response Object
type AddDeviceResponse struct {

	// 设备ID
	DeviceId *string `json:"device_id,omitempty"`

	// 设备密钥，认证类型使用密钥认证接入(SECRET)可填写该字段。注意：NB设备密钥由于协议特殊性，只支持十六进制密钥接入；修改设备、查询设备及查询设备列表接口不返回该参数。
	Password       *string `json:"password,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddDeviceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDeviceResponse struct{}"
	}

	return strings.Join([]string{"AddDeviceResponse", string(data)}, " ")
}
