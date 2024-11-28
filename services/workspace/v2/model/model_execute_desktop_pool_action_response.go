package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteDesktopPoolActionResponse Response Object
type ExecuteDesktopPoolActionResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteDesktopPoolActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDesktopPoolActionResponse struct{}"
	}

	return strings.Join([]string{"ExecuteDesktopPoolActionResponse", string(data)}, " ")
}
