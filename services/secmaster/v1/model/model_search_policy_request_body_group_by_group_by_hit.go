package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchPolicyRequestBodyGroupByGroupByHit 聚合结果
type SearchPolicyRequestBodyGroupByGroupByHit struct {

	// 源字段
	Source *string `json:"source,omitempty"`

	// 目标字段
	Dest *string `json:"dest,omitempty"`
}

func (o SearchPolicyRequestBodyGroupByGroupByHit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchPolicyRequestBodyGroupByGroupByHit struct{}"
	}

	return strings.Join([]string{"SearchPolicyRequestBodyGroupByGroupByHit", string(data)}, " ")
}
