package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubnetTagsRequest Request Object
type ListSubnetTagsRequest struct {
}

func (o ListSubnetTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubnetTagsRequest struct{}"
	}

	return strings.Join([]string{"ListSubnetTagsRequest", string(data)}, " ")
}
