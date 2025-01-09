package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesktopSidsResponse Response Object
type UpdateDesktopSidsResponse struct {

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDesktopSidsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesktopSidsResponse struct{}"
	}

	return strings.Join([]string{"UpdateDesktopSidsResponse", string(data)}, " ")
}
