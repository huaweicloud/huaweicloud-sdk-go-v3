package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartProtectionGroupResponse Response Object
type StartProtectionGroupResponse struct {

	// 成功返回jobId信息
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartProtectionGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartProtectionGroupResponse struct{}"
	}

	return strings.Join([]string{"StartProtectionGroupResponse", string(data)}, " ")
}
