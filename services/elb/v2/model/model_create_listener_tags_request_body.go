package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type CreateListenerTagsRequestBody struct {
	Tag *ResourceTag `json:"tag"`
}

func (o CreateListenerTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerTagsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateListenerTagsRequestBody", string(data)}, " ")
}
