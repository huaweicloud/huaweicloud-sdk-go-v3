package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRetryJobResponse Response Object
type BatchRetryJobResponse struct {
	Body           *[]BatchOperateJobRsp `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o BatchRetryJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRetryJobResponse struct{}"
	}

	return strings.Join([]string{"BatchRetryJobResponse", string(data)}, " ")
}
