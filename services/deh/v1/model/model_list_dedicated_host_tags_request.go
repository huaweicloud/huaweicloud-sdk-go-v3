package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDedicatedHostTagsRequest Request Object
type ListDedicatedHostTagsRequest struct {
}

func (o ListDedicatedHostTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDedicatedHostTagsRequest struct{}"
	}

	return strings.Join([]string{"ListDedicatedHostTagsRequest", string(data)}, " ")
}
