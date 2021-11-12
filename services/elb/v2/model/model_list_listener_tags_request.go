package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListListenerTagsRequest struct {
}

func (o ListListenerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListListenerTagsRequest struct{}"
	}

	return strings.Join([]string{"ListListenerTagsRequest", string(data)}, " ")
}
