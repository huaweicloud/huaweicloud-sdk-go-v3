package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuraSessionStatementRecordsResponse Response Object
type ListAuraSessionStatementRecordsResponse struct {

	// **参数解释**：本次查询记录。 **取值范围**：不涉及。
	Records *[]SessionStatementRecord `json:"records,omitempty"`

	// **参数解释**：本次查询记录总数。 **取值范围**：1~2147483647。
	CurrentCount *int32 `json:"current_count,omitempty"`

	// **参数解释**：下一页查询marker值，若为空表示当前已是最后一页。 **取值范围**：长度为1~64个字符，支持大小写英文字母、数字、连字符。
	NextMarker     *string `json:"next_marker,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAuraSessionStatementRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuraSessionStatementRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListAuraSessionStatementRecordsResponse", string(data)}, " ")
}
