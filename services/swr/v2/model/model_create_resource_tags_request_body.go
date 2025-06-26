package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateResourceTagsRequestBody struct {

	// 批量操作标签
	Tags []ResourceTag `json:"tags"`
}

func (o CreateResourceTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceTagsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateResourceTagsRequestBody", string(data)}, " ")
}
