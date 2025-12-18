package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceLbResponse Response Object
type UpdateInstanceLbResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateInstanceLbResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceLbResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceLbResponse", string(data)}, " ")
}
