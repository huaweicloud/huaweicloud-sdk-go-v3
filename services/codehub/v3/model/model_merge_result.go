package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MergeResult struct {

	// 已关闭合并请求数
	Closed *float64 `json:"closed,omitempty"`

	// 合并请求详情
	MergeRequests *[]MergeRequestsItem `json:"merge_requests,omitempty"`

	// 已合并合并请求数数
	Merged *float64 `json:"merged,omitempty"`

	// 开启中合并请求数
	Opened *float64 `json:"opened,omitempty"`

	// 合并请求总数
	Total *float64 `json:"total,omitempty"`
}

func (o MergeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeResult struct{}"
	}

	return strings.Join([]string{"MergeResult", string(data)}, " ")
}
