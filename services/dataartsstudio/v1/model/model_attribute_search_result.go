package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttributeSearchResult 请求参数
type AttributeSearchResult struct {

	// 名称列表
	Name *[]string `json:"name,omitempty"`

	// 值列表
	Values *[]interface{} `json:"values,omitempty"`
}

func (o AttributeSearchResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttributeSearchResult struct{}"
	}

	return strings.Join([]string{"AttributeSearchResult", string(data)}, " ")
}
