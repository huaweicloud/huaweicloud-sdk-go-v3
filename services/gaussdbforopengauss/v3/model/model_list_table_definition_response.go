package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTableDefinitionResponse Response Object
type ListTableDefinitionResponse struct {

	// **参数解释**: 数据库表定义信息。 **取值范围**: 不涉及。
	TableDefinition *string `json:"table_definition,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o ListTableDefinitionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableDefinitionResponse struct{}"
	}

	return strings.Join([]string{"ListTableDefinitionResponse", string(data)}, " ")
}
