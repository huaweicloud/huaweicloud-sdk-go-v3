package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceProjectTagsRequest Request Object
type ListInstanceProjectTagsRequest struct {
}

func (o ListInstanceProjectTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceProjectTagsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceProjectTagsRequest", string(data)}, " ")
}
