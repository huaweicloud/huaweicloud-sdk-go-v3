package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKillProcessTaskPreviewResponse Response Object
type ShowKillProcessTaskPreviewResponse struct {

	// **参数解释**：  会话总数，最多支持预览20条会话记录。  **取值范围**：  0-20。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**：  会话记录。
	Processes      *[]ProcessDetail `json:"processes,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowKillProcessTaskPreviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKillProcessTaskPreviewResponse struct{}"
	}

	return strings.Join([]string{"ShowKillProcessTaskPreviewResponse", string(data)}, " ")
}
