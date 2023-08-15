package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourceTagRequestBody 添加资源标签的请求体。
type CreateResourceTagRequestBody struct {
	Tag *Tag `json:"tag"`
}

func (o CreateResourceTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceTagRequestBody struct{}"
	}

	return strings.Join([]string{"CreateResourceTagRequestBody", string(data)}, " ")
}
