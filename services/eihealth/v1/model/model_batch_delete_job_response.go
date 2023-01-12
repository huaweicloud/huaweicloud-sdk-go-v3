package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteJobResponse struct {
	Body           *[]BatchOperateJobRsp `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o BatchDeleteJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteJobResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteJobResponse", string(data)}, " ")
}
