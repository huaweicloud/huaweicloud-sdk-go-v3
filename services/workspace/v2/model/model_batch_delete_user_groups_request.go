package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteUserGroupsRequest Request Object
type BatchDeleteUserGroupsRequest struct {
	Body *BatchDeleteUserGroupsReq `json:"body,omitempty"`
}

func (o BatchDeleteUserGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteUserGroupsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteUserGroupsRequest", string(data)}, " ")
}
