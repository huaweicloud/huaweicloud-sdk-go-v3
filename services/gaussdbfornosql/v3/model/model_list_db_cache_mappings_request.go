package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDbCacheMappingsRequest Request Object
type ListDbCacheMappingsRequest struct {

	// 映射ID，可以调用“查询内存加速映射列表和详情”接口获取。
	Id *string `json:"id,omitempty"`

	// 映射名称。名称以“*”起始，表示按照“*”后面的值模糊匹配，否则，按照实际填写的名称精确匹配查询。
	Name *string `json:"name,omitempty"`

	// 源实例ID。
	SourceInstanceId *string `json:"source_instance_id,omitempty"`

	// 源实例名称。名称以“*”起始，表示按照“*”后面的值模糊匹配，否则，按照实际填写的名称精确匹配查询。
	SourceInstanceName *string `json:"source_instance_name,omitempty"`

	// 目标实例ID。
	TargetInstanceId *string `json:"target_instance_id,omitempty"`

	// 目标实例名称。名称以“*”起始，表示按照“*”后面的值模糊匹配，否则，按照实际填写的名称精确匹配查询。
	TargetInstanceName *string `json:"target_instance_name,omitempty"`

	// 索引位置，偏移量。 从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询）。 取值必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询个数上限值。取值范围：1~100。不传该参数时，默认查询前100条信息。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDbCacheMappingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbCacheMappingsRequest struct{}"
	}

	return strings.Join([]string{"ListDbCacheMappingsRequest", string(data)}, " ")
}
