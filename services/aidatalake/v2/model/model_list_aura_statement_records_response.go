package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuraStatementRecordsResponse Response Object
type ListAuraStatementRecordsResponse struct {

	// **参数解释**：本次查询记录总数。 **取值范围**：1~100。
	CurrentCount *int32 `json:"current_count,omitempty"`

	// **参数解释**：下一页查询marker值，若为空表示当前已是最后一页。 **取值范围**：长度为1~64个字符，支持大小写英文字母、数字、连字符。
	NextMarker *string `json:"next_marker,omitempty"`

	// **参数解释**：sql历史记录。
	Records        *[]StatementRecord `json:"records,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListAuraStatementRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuraStatementRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListAuraStatementRecordsResponse", string(data)}, " ")
}
