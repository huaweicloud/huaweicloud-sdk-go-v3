package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SearchPolicyRequestBodySort struct {

	// 排序字段
	SortBy *string `json:"sort_by,omitempty"`

	// 顺序或倒序
	Order *string `json:"order,omitempty"`
}

func (o SearchPolicyRequestBodySort) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchPolicyRequestBodySort struct{}"
	}

	return strings.Join([]string{"SearchPolicyRequestBodySort", string(data)}, " ")
}
