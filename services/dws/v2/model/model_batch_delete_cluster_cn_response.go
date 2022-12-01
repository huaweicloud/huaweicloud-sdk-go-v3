package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteClusterCnResponse struct {

	// 批量删除CN节点任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteClusterCnResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteClusterCnResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteClusterCnResponse", string(data)}, " ")
}
