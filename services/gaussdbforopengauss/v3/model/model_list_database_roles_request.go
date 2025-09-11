package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseRolesRequest Request Object
type ListDatabaseRolesRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// 偏移量表示从此偏移量开始查询, offset大于等于0，默认0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量,取值范围[1, 100]，默认10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDatabaseRolesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseRolesRequest struct{}"
	}

	return strings.Join([]string{"ListDatabaseRolesRequest", string(data)}, " ")
}
