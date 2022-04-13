package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPublicipTagsRequest struct {
}

func (o ListPublicipTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicipTagsRequest struct{}"
	}

	return strings.Join([]string{"ListPublicipTagsRequest", string(data)}, " ")
}
