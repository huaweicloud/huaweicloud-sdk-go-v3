package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateUsersRequest Request Object
type BatchCreateUsersRequest struct {
	Body *BatchCreateUsersReq `json:"body,omitempty"`
}

func (o BatchCreateUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateUsersRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateUsersRequest", string(data)}, " ")
}
