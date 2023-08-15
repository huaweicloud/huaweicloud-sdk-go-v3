package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchOperateTagRequestBody struct {

	// 标签列表，key和value键值对的集合。
	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o BatchOperateTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOperateTagRequestBody struct{}"
	}

	return strings.Join([]string{"BatchOperateTagRequestBody", string(data)}, " ")
}
