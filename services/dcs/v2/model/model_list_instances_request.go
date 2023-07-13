package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesRequest Request Object
type ListInstancesRequest struct {

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 是否返回创建失败的实例数。 当参数值为“true”时，返回创建失败的实例数。参数值为“false”或者其他值，不返回创建失败的实例数。
	IncludeFailure *string `json:"include_failure,omitempty"`

	// 是否返回已删除的实例数。 当参数值为“true”时，返回已删除的实例数。参数值为“false”或者其他值，不返回已删除的实例数。
	IncludeDelete *string `json:"include_delete,omitempty"`

	// 实例名称。
	Name *string `json:"name,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示条数，最小值为1，最大值为1000，若不设置该参数，则为10。
	Limit *int32 `json:"limit,omitempty"`

	// 实例状态。详细状态说明见[缓存实例状态说明](https://support.huaweicloud.com/api-dcs/dcs-api-0312047.html)
	Status *string `json:"status,omitempty"`

	// 是否按照实例名称进行精确匹配查询。  和name字段对应，name字段为模糊匹配的用例名，name_equal是精确匹配的用例名。
	NameEqual *string `json:"name_equal,omitempty"`

	// 根据实例标签键值对进行查询。{key}表示标签键，{value}表示标签值。  如果同时使用多个标签键值对进行查询，中间使用逗号分隔开，表示查询同时包含指定标签键值对的实例。
	Tags *string `json:"tags,omitempty"`

	// 连接缓存实例的IP地址。
	Ip *string `json:"ip,omitempty"`
}

func (o ListInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesRequest", string(data)}, " ")
}
