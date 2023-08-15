package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateTagsRequestBody struct {

	// 操作标识： - create，创建
	Action string `json:"action"`

	// 标签列表。
	Tags []Tag `json:"tags"`
}

func (o BatchCreateTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateTagsRequestBody", string(data)}, " ")
}
