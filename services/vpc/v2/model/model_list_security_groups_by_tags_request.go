package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityGroupsByTagsRequest Request Object
type ListSecurityGroupsByTagsRequest struct {
	Body *ListSecurityGroupsByTagsRequestBody `json:"body,omitempty"`
}

func (o ListSecurityGroupsByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityGroupsByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityGroupsByTagsRequest", string(data)}, " ")
}
