package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPasswordlessConfigRequest Request Object
type ShowPasswordlessConfigRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 索引位置偏移量，表示从查询到的免密配置列表偏移offset条数据后查询对应的免密配置。 取值大于或等于0。不传该参数时，查询偏移量默认为0，表示从第一条记录开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 查询个数上限值。 取值范围：1~100。不传该参数时，默认查询前100条记录。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowPasswordlessConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPasswordlessConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowPasswordlessConfigRequest", string(data)}, " ")
}
