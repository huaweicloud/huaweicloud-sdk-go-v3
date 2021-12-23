package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchUpdateMembersRequest struct {
	Body *BatchUpdateMembersRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateMembersRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateMembersRequest", string(data)}, " ")
}
