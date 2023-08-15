package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryDcDeviceRespDto 查询数采设备响应结构体
type QueryDcDeviceRespDto struct {

	// 设备id
	DeviceId *string `json:"device_id,omitempty"`
}

func (o QueryDcDeviceRespDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDcDeviceRespDto struct{}"
	}

	return strings.Join([]string{"QueryDcDeviceRespDto", string(data)}, " ")
}
