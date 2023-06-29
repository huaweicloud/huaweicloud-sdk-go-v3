package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteResourceTags 标签列表信息
type BatchDeleteResourceTags struct {

	// 标签列表
	Tags []BatchDeleteResourceTag `json:"tags"`
}

func (o BatchDeleteResourceTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteResourceTags struct{}"
	}

	return strings.Join([]string{"BatchDeleteResourceTags", string(data)}, " ")
}
