package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreatePortTagsRequestBody This is a auto create Body Object
type BatchCreatePortTagsRequestBody struct {

	// 标签列表 约束：最大支持20组标签键值对
	Tags []ResourceTag `json:"tags"`
}

func (o BatchCreatePortTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreatePortTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreatePortTagsRequestBody", string(data)}, " ")
}
