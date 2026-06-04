package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvalibleDdmsResponse Response Object
type ShowAvalibleDdmsResponse struct {

	// 可用目标DDM。
	Instances *[]DdmInstance4Restore `json:"instances,omitempty"`

	// **参数解释**：  分页参数: 起始值。  **取值范围**：   大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：  分页参数: 每页记录数。  **取值范围**：  大于0且小于等于128。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：  总记录数。  **取值范围**：  不涉及。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAvalibleDdmsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvalibleDdmsResponse struct{}"
	}

	return strings.Join([]string{"ShowAvalibleDdmsResponse", string(data)}, " ")
}
