package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateTagsRequestBody struct {

	// 标签列表
	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o BatchUpdateTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateTagsRequestBody", string(data)}, " ")
}
