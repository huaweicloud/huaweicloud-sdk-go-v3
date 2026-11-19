package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSqlPatchRequestBody struct {

	// **参数解释**: 实例ID，与路径参数instance_id一致。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	QueryInfo *SqlPatchQueryInfoOption `json:"query_info,omitempty"`

	// **参数解释**: 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。 **约束限制**: 不涉及。 **取值范围**: 大于等于0的整数。 **默认取值**: 0
	Offset int32 `json:"offset"`

	// **参数解释**: 查询记录数。 **约束限制**: 不能为负数。 **取值范围**: 最小值为1，最大值为1000。 **默认取值**: 20
	Limit int32 `json:"limit"`
}

func (o ListSqlPatchRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlPatchRequestBody struct{}"
	}

	return strings.Join([]string{"ListSqlPatchRequestBody", string(data)}, " ")
}
