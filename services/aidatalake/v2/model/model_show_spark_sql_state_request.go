package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSparkSqlStateRequest Request Object
type ShowSparkSqlStateRequest struct {

	// **参数解释**：工作空间ID，可通过[获取工作空间ID](aidatalake_02_0044.xml)获取。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：statement id。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、下划线。 **默认取值**：不涉及。
	StatementId string `json:"statement_id"`
}

func (o ShowSparkSqlStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkSqlStateRequest struct{}"
	}

	return strings.Join([]string{"ShowSparkSqlStateRequest", string(data)}, " ")
}
