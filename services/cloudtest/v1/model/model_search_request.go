package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SearchRequest struct {

	// 过滤类型
	SearchType *string `json:"search_type,omitempty"`

	// 过滤条件
	SearchValue *string `json:"search_value,omitempty"`
}

func (o SearchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchRequest struct{}"
	}

	return strings.Join([]string{"SearchRequest", string(data)}, " ")
}
