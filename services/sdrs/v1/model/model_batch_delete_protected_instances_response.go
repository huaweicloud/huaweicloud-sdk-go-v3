package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteProtectedInstancesResponse Response Object
type BatchDeleteProtectedInstancesResponse struct {

	// 成功返回jobId信息
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteProtectedInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteProtectedInstancesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteProtectedInstancesResponse", string(data)}, " ")
}
