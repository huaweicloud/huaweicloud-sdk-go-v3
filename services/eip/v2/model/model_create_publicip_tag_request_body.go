package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建tag对象的请求体
type CreatePublicipTagRequestBody struct {
	Tag *ResourceTagOption `json:"tag"`
}

func (o CreatePublicipTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicipTagRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePublicipTagRequestBody", string(data)}, " ")
}
