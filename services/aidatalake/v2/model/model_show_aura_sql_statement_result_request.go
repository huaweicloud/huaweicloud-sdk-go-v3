package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuraSqlStatementResultRequest Request Object
type ShowAuraSqlStatementResultRequest struct {

	// **参数解释**：工作空间的ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：会话ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	SessionId string `json:"session_id"`

	// **参数解释**：statement id。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、下划线。 **默认取值**：不涉及。
	StatementId string `json:"statement_id"`

	// **参数解释**：查询页码。 **约束限制**：不涉及。 **取值范围**：1~2147483647。 **默认取值**：不涉及。
	PageNum *int32 `json:"page_num,omitempty"`

	// **参数解释**：是否返回结果集的obs路径。 **约束限制**：不涉及。 **取值范围**：   - true：返回结果集的obs路径。   - false：返回结果集。 **默认取值**：不涉及。
	IsEnableObsPath *bool `json:"is_enable_obs_path,omitempty"`
}

func (o ShowAuraSqlStatementResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuraSqlStatementResultRequest struct{}"
	}

	return strings.Join([]string{"ShowAuraSqlStatementResultRequest", string(data)}, " ")
}
