package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchCreateClusterCnResponse struct {

	// 批量增加CN节点任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchCreateClusterCnResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateClusterCnResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateClusterCnResponse", string(data)}, " ")
}
