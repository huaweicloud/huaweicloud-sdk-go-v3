package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteNodesResponse Response Object
type BatchDeleteNodesResponse struct {

	// 任务id
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteNodesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteNodesResponse", string(data)}, " ")
}
