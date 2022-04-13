package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标签
type CreateSharedTagRequestBody struct {
	Tag *ResourceTag `json:"tag"`
}

func (o CreateSharedTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSharedTagRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSharedTagRequestBody", string(data)}, " ")
}
