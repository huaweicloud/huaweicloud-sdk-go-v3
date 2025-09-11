package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDatabaseSchemaRequest Request Object
type DeleteDatabaseSchemaRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *DeleteDatabaseSchemaRequestBody `json:"body,omitempty"`
}

func (o DeleteDatabaseSchemaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseSchemaRequest struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseSchemaRequest", string(data)}, " ")
}
