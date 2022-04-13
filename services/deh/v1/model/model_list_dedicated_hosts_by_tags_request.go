package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDedicatedHostsByTagsRequest struct {
	Body *ReqListDehByTags `json:"body,omitempty"`
}

func (o ListDedicatedHostsByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDedicatedHostsByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListDedicatedHostsByTagsRequest", string(data)}, " ")
}
