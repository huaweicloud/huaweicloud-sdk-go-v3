package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchOverDisasterRecoveryResponse Response Object
type SwitchOverDisasterRecoveryResponse struct {

	// **参数解释：** 容灾倒换任务ID。 **取值范围：** 不涉及
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchOverDisasterRecoveryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchOverDisasterRecoveryResponse struct{}"
	}

	return strings.Join([]string{"SwitchOverDisasterRecoveryResponse", string(data)}, " ")
}
