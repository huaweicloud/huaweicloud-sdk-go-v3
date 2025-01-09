package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccessModeResponse Response Object
type UpdateAccessModeResponse struct {

	// 任务id
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAccessModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessModeResponse struct{}"
	}

	return strings.Join([]string{"UpdateAccessModeResponse", string(data)}, " ")
}
