package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteTagsRequestBody struct {

	// 操作标识： - delete，刪除
	Action string `json:"action"`

	// 标签列表。
	Tags []Tag `json:"tags"`
}

func (o BatchDeleteTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteTagsRequestBody", string(data)}, " ")
}
