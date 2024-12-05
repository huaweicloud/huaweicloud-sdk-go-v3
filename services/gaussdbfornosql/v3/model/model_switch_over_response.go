package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchOverResponse Response Object
type SwitchOverResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchOverResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchOverResponse struct{}"
	}

	return strings.Join([]string{"SwitchOverResponse", string(data)}, " ")
}
