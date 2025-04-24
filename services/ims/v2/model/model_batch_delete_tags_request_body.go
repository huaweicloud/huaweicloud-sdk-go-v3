package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteTagsRequestBody 批量删除镜像标签请求体
type BatchDeleteTagsRequestBody struct {

	// 需要删除的标签键值对集合。
	Tags []ResourceTag `json:"tags"`
}

func (o BatchDeleteTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteTagsRequestBody", string(data)}, " ")
}
