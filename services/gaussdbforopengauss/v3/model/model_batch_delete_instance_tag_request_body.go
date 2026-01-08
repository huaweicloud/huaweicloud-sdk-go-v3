package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteInstanceTagRequestBody struct {

	// **参数解释**: 需要删除的标签列表。
	Keys []string `json:"keys"`
}

func (o BatchDeleteInstanceTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstanceTagRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstanceTagRequestBody", string(data)}, " ")
}
