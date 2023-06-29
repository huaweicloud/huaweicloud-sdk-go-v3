package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllTagsRequest Request Object
type ListAllTagsRequest struct {
}

func (o ListAllTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllTagsRequest struct{}"
	}

	return strings.Join([]string{"ListAllTagsRequest", string(data)}, " ")
}
