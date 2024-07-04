package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStarrocksParamsRequest Request Object
type ShowStarrocksParamsRequest struct {

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 节点类型。取值范围: - be - fe
	NodeType string `json:"node_type"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowStarrocksParamsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStarrocksParamsRequest struct{}"
	}

	return strings.Join([]string{"ShowStarrocksParamsRequest", string(data)}, " ")
}
