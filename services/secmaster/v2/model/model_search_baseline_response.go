package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchBaselineResponse Response Object
type SearchBaselineResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 查询结果总数
	Total *int32 `json:"total,omitempty"`

	// 分页大小
	Size *int32 `json:"size,omitempty"`

	// 偏移量
	Page *int32 `json:"page,omitempty"`

	// 是否成功
	Success *bool `json:"success,omitempty"`

	// 查询结果列表
	Data           *[]string `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o SearchBaselineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchBaselineResponse struct{}"
	}

	return strings.Join([]string{"SearchBaselineResponse", string(data)}, " ")
}
