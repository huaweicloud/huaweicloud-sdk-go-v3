package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRegisterDefenderPluginResponse Response Object
type BatchRegisterDefenderPluginResponse struct {

	// **参数解释**: 任务ID **取值范围**: 字符长度1-128位
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2000000
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**: 接受数量 **取值范围**: 最小值0，最大值2000000
	AcceptedCount  *int32 `json:"accepted_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchRegisterDefenderPluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRegisterDefenderPluginResponse struct{}"
	}

	return strings.Join([]string{"BatchRegisterDefenderPluginResponse", string(data)}, " ")
}
