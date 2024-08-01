package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExternalAccessDetails 外部访问分析详细结果。
type ExternalAccessDetails struct {

	// 允许外部主体使用的操作。
	Action []string `json:"action"`

	// 分析的策略语句中导致访问分析结果的条件。
	Condition []FindingCondition `json:"condition"`

	// 表示生成访问分析结果的策略是否允许公共访问资源。
	IsPublic bool `json:"is_public"`

	Principal *FindingPrincipal `json:"principal"`

	// 访问分析结果的来源，这指示如何授予生成访问分析结果的访问权限。
	Sources *[]FindingSourceType `json:"sources,omitempty"`
}

func (o ExternalAccessDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalAccessDetails struct{}"
	}

	return strings.Join([]string{"ExternalAccessDetails", string(data)}, " ")
}
