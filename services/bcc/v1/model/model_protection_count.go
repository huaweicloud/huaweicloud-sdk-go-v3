package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProtectionCount 未保护资源统计
type ProtectionCount struct {

	// 总数
	Total int32 `json:"total"`

	// 未定级资源总数
	Ungraded *int32 `json:"ungraded,omitempty"`

	Graded *GradedCount `json:"graded,omitempty"`
}

func (o ProtectionCount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectionCount struct{}"
	}

	return strings.Join([]string{"ProtectionCount", string(data)}, " ")
}
