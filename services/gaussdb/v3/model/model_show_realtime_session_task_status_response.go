package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRealtimeSessionTaskStatusResponse Response Object
type ShowRealtimeSessionTaskStatusResponse struct {

	// **参数解释**：  收集全部实时会话信息任务状态。  **取值范围**：  - running: 运行中。 - finished：已完成。 - error: 发生错误。 - unrecognized：未识别状态。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRealtimeSessionTaskStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRealtimeSessionTaskStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowRealtimeSessionTaskStatusResponse", string(data)}, " ")
}
