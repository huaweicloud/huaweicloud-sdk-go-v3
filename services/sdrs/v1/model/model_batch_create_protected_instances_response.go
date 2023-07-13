package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateProtectedInstancesResponse Response Object
type BatchCreateProtectedInstancesResponse struct {

	// 成功返回jobId信息
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchCreateProtectedInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateProtectedInstancesResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateProtectedInstancesResponse", string(data)}, " ")
}
