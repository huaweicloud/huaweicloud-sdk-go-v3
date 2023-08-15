package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityGroupResponse Response Object
type UpdateSecurityGroupResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateSecurityGroupResponse", string(data)}, " ")
}
