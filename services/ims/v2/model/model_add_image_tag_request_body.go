package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddImageTagRequestBody 请求参数
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
