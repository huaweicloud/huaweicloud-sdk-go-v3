package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotifyReplaceNodeResponse Response Object
type NotifyReplaceNodeResponse struct {

	// 任务流id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o NotifyReplaceNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotifyReplaceNodeResponse struct{}"
	}

	return strings.Join([]string{"NotifyReplaceNodeResponse", string(data)}, " ")
}
