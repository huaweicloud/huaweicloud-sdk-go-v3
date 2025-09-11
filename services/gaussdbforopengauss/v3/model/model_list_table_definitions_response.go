package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTableDefinitionsResponse Response Object
type ListTableDefinitionsResponse struct {

	// **参数解释**: 数据库表定义信息列表。
	TableDefinitions *[]SchemaTableDefinitionResult `json:"table_definitions,omitempty"`
	HttpStatusCode   int                            `json:"-"`
}

func (o ListTableDefinitionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableDefinitionsResponse struct{}"
	}

	return strings.Join([]string{"ListTableDefinitionsResponse", string(data)}, " ")
}
