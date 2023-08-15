package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTagsRequestBody create tag request
type CreateTagsRequestBody struct {

	// 标签列表。
	Tags []ResourceTag `json:"tags"`
}

func (o CreateTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTagsRequestBody", string(data)}, " ")
}
