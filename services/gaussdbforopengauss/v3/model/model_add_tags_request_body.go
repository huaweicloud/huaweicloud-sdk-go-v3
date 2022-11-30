package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddTagsRequestBody struct {
	Tags []TagsOption `json:"tags"`
}

func (o AddTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTagsRequestBody struct{}"
	}

	return strings.Join([]string{"AddTagsRequestBody", string(data)}, " ")
}
