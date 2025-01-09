package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDdmConfigurationsRequest Request Object
type ListDdmConfigurationsRequest struct {

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0。取值必须为数字，且不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询个数上限值。取值范围：1~128。不传该参数时，默认值为10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDdmConfigurationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDdmConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"ListDdmConfigurationsRequest", string(data)}, " ")
}
