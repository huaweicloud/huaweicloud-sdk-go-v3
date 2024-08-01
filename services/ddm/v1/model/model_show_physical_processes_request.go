package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPhysicalProcessesRequest Request Object
type ShowPhysicalProcessesRequest struct {

	// 关联RDS的ID。
	InstanceId string `json:"instance_id"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0。取值必须为数字，且不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询个数上限值。取值范围：1~128。不传该参数时，默认值为10。
	Limit *int32 `json:"limit,omitempty"`

	// 会话结果筛选关键子，长度最大255
	Keyword *string `json:"keyword,omitempty"`
}

func (o ShowPhysicalProcessesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPhysicalProcessesRequest struct{}"
	}

	return strings.Join([]string{"ShowPhysicalProcessesRequest", string(data)}, " ")
}
