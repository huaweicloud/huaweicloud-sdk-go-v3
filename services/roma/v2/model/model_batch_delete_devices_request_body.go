package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteDevicesRequestBody struct {
	// 删除设备ID列表，自动向下取整

	Resources []int32 `json:"resources"`
}

func (o BatchDeleteDevicesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDevicesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteDevicesRequestBody", string(data)}, " ")
}
