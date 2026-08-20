package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ThirdPartyAssociatedResult 工作项关联外部链接查询结果
type ThirdPartyAssociatedResult struct {
	Data *ThirdPartyAssociatedResultData `json:"data,omitempty"`

	// 工作项关联外部链接总数。
	Count *string `json:"count,omitempty"`
}

func (o ThirdPartyAssociatedResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThirdPartyAssociatedResult struct{}"
	}

	return strings.Join([]string{"ThirdPartyAssociatedResult", string(data)}, " ")
}
