package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNatTagRequestBody 请求参数。
type CreateNatTagRequestBody struct {
	Tag *TagBody `json:"tag"`
}

func (o CreateNatTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatTagRequestBody struct{}"
	}

	return strings.Join([]string{"CreateNatTagRequestBody", string(data)}, " ")
}
