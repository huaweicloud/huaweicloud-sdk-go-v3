package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListConfigurationTemplatesRequest struct {
	// 索引位置，偏移量。   - 从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询）。   - 取值必须为数字，不能为负数。

	Offset *int32 `json:"offset,omitempty"`
	// 查询个数上限值。 - 取值范围: 1~100。 - 不传该参数时，默认查询前100条信息。

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListConfigurationTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListConfigurationTemplatesRequest", string(data)}, " ")
}
