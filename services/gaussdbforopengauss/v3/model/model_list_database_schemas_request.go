package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseSchemasRequest Request Object
type ListDatabaseSchemasRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// 数据库名称。
	DbName string `json:"db_name"`

	// 偏移量表示从此偏移量开始查询, offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量,取值范围[1, 100]。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDatabaseSchemasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseSchemasRequest struct{}"
	}

	return strings.Join([]string{"ListDatabaseSchemasRequest", string(data)}, " ")
}
