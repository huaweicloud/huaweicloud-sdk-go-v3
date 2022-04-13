package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDevicesRequest struct {
	// 查询的起始位置，取值范围为非负整数，默认为0

	Offset *int32 `json:"offset,omitempty"`
	// 每页显示的条目数量，取值范围1~100，默认为100

	Limit *int32 `json:"limit,omitempty"`
	// 设备名称，模糊匹配，只允许中英文、数字、下划线、中划线，长度1-60

	Name *string `json:"name,omitempty"`
}

func (o ListDevicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDevicesRequest struct{}"
	}

	return strings.Join([]string{"ListDevicesRequest", string(data)}, " ")
}
