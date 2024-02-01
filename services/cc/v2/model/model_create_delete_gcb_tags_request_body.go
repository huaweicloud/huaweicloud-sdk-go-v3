package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDeleteGcbTagsRequestBody 资源标签Tag汇总复数。
type CreateDeleteGcbTagsRequestBody struct {

	// 批量添加/删除资源标签。
	Tags []RequiredTag `json:"tags"`
}

func (o CreateDeleteGcbTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeleteGcbTagsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDeleteGcbTagsRequestBody", string(data)}, " ")
}
