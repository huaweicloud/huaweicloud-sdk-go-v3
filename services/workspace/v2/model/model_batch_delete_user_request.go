package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteUserRequest Request Object
type BatchDeleteUserRequest struct {
	Body *BatchDeleteUserReq `json:"body,omitempty"`
}

func (o BatchDeleteUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteUserRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteUserRequest", string(data)}, " ")
}
