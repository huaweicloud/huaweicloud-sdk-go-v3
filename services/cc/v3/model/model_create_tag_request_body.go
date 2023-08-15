package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTagRequestBody 创建Tag请求体
type CreateTagRequestBody struct {
	Tag *Tag `json:"tag,omitempty"`
}

func (o CreateTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTagRequestBody", string(data)}, " ")
}
