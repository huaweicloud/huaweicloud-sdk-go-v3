package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DriverInfo struct {

	// 文件名称。
	DriverName *string `json:"driver_name,omitempty"`

	// 最后修改时间。
	LastModified *string `json:"last_modified,omitempty"`

	// 文件大小，单位：byte
	Size *int32 `json:"size,omitempty"`
}

func (o DriverInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DriverInfo struct{}"
	}

	return strings.Join([]string{"DriverInfo", string(data)}, " ")
}
