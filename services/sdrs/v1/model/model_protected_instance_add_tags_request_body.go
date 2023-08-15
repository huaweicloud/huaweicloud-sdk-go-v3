package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProtectedInstanceAddTagsRequestBody 添加保护实例标签请求体
type ProtectedInstanceAddTagsRequestBody struct {
	Tag *ResourceTag `json:"tag"`
}

func (o ProtectedInstanceAddTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectedInstanceAddTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ProtectedInstanceAddTagsRequestBody", string(data)}, " ")
}
