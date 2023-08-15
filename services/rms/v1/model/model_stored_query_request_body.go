package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StoredQueryRequestBody struct {

	// ResourceQL 名字
	Name string `json:"name"`

	// ResourceQL 描述
	Description *string `json:"description,omitempty"`

	// ResourceQL 表达式
	Expression string `json:"expression"`
}

func (o StoredQueryRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StoredQueryRequestBody struct{}"
	}

	return strings.Join([]string{"StoredQueryRequestBody", string(data)}, " ")
}
