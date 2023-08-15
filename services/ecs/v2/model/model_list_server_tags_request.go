package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServerTagsRequest Request Object
type ListServerTagsRequest struct {
}

func (o ListServerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerTagsRequest struct{}"
	}

	return strings.Join([]string{"ListServerTagsRequest", string(data)}, " ")
}
