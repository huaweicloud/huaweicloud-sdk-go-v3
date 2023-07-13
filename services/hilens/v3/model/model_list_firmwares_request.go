package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFirmwaresRequest Request Object
type ListFirmwaresRequest struct {

	// 固件适用设备类型
	DeviceType string `json:"device_type"`

	// 边缘节点架构
	Arch *string `json:"arch,omitempty"`

	// 边缘设备操作系统名称
	OsName *string `json:"os_name,omitempty"`

	// 边缘设备操作系统版本
	OsVersion *string `json:"os_version,omitempty"`

	// 查询的起始位置，取值范围为非负整数，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量，取值范围1~100，默认为100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListFirmwaresRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFirmwaresRequest struct{}"
	}

	return strings.Join([]string{"ListFirmwaresRequest", string(data)}, " ")
}
