package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateTopicAttributeRequestBody struct {
	// 1. 当请求参数name为access_policy时，value为Topic属性值，最大支持30KB。  2. 当请求参数name为introduction时，value为topic简介，最大长度120B。

	Value string `json:"value"`
}

func (o UpdateTopicAttributeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicAttributeRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTopicAttributeRequestBody", string(data)}, " ")
}
