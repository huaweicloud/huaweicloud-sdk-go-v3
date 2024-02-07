package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateResourcesTagsRequestBody struct {
	Tags []ResourceTag `json:"tags"`
}

func (o CreateResourcesTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourcesTagsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateResourcesTagsRequestBody", string(data)}, " ")
}
