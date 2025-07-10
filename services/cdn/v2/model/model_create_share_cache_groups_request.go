package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateShareCacheGroupsRequest Request Object
type CreateShareCacheGroupsRequest struct {
	Body *CreateShareCacheGroupsRequstBody `json:"body,omitempty"`
}

func (o CreateShareCacheGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShareCacheGroupsRequest struct{}"
	}

	return strings.Join([]string{"CreateShareCacheGroupsRequest", string(data)}, " ")
}
