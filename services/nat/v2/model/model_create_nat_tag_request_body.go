package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNatTagRequestBody 请求参数。
type CreateNatTagRequestBody struct {

	// 标签列表。请参考表TagBody字段数据结构说明。
	Tag *interface{} `json:"tag"`
}

func (o CreateNatTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatTagRequestBody struct{}"
	}

	return strings.Join([]string{"CreateNatTagRequestBody", string(data)}, " ")
}
