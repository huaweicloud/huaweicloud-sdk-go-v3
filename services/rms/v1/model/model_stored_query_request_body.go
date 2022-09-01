package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StoredQueryRequestBody struct {

	// ResourceQL 名字
	Name string `json:"name" xml:"name"`

	// ResourceQL 描述
	Description *string `json:"description,omitempty" xml:"description"`

	// ResourceQL 表达式
	Expression string `json:"expression" xml:"expression"`
}

func (o StoredQueryRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StoredQueryRequestBody struct{}"
	}

	return strings.Join([]string{"StoredQueryRequestBody", string(data)}, " ")
}
