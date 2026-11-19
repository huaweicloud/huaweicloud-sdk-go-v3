package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaurusDbSlowProcessThresholdResponse Response Object
type ShowTaurusDbSlowProcessThresholdResponse struct {

	// **参数解释**:  慢会话阈值，单位为秒。  **取值范围**：  1-86400。
	SlowProcessThreshold *int32 `json:"slow_process_threshold,omitempty"`
	HttpStatusCode       int    `json:"-"`
}

func (o ShowTaurusDbSlowProcessThresholdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaurusDbSlowProcessThresholdResponse struct{}"
	}

	return strings.Join([]string{"ShowTaurusDbSlowProcessThresholdResponse", string(data)}, " ")
}
