package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryRunRequestBody struct {

	// ResourceQL 表达式
	Expression string `json:"expression"`
}

func (o QueryRunRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryRunRequestBody struct{}"
	}

	return strings.Join([]string{"QueryRunRequestBody", string(data)}, " ")
}
