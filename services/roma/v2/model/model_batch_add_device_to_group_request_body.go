package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchAddDeviceToGroupRequestBody struct {
	// 设备ID列表，自动向下取整

	Resources []int32 `json:"resources"`
}

func (o BatchAddDeviceToGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddDeviceToGroupRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAddDeviceToGroupRequestBody", string(data)}, " ")
}
