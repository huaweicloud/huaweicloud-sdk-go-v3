package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 请求参数
type AddImageTagRequestBody struct {
	Tag *ResourceTag `json:"tag"`
}

func (o AddImageTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddImageTagRequestBody struct{}"
	}

	return strings.Join([]string{"AddImageTagRequestBody", string(data)}, " ")
}
