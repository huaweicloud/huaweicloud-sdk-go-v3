package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchCancelJobResponse struct {
	Body           *[]BatchOperateJobRsp `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o BatchCancelJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCancelJobResponse struct{}"
	}

	return strings.Join([]string{"BatchCancelJobResponse", string(data)}, " ")
}
