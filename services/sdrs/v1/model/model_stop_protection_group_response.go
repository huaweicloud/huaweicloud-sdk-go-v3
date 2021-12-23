package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StopProtectionGroupResponse struct {
	// 成功返回jobId信息

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopProtectionGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopProtectionGroupResponse struct{}"
	}

	return strings.Join([]string{"StopProtectionGroupResponse", string(data)}, " ")
}
