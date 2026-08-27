package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteEvolveTaskRequest Request Object
type BatchDeleteEvolveTaskRequest struct {
	Body *EvolveTaskBatchDeleteReq `json:"body,omitempty"`
}

func (o BatchDeleteEvolveTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteEvolveTaskRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteEvolveTaskRequest", string(data)}, " ")
}
