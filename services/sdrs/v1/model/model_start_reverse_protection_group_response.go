package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StartReverseProtectionGroupResponse struct {

	// 成功返回jobId信息
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartReverseProtectionGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartReverseProtectionGroupResponse struct{}"
	}

	return strings.Join([]string{"StartReverseProtectionGroupResponse", string(data)}, " ")
}
