package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// id, 领域 14, '性能', 15, '功能', 16, '可靠性' 17, '网络安全' 18, '可维护性' 19, '其他DFX' 20, '可用性'
type IssueItemSfV4Domain struct {
	// 领域id

	Id *int32 `json:"id,omitempty"`
	// 领域

	Name *string `json:"name,omitempty"`
}

func (o IssueItemSfV4Domain) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueItemSfV4Domain struct{}"
	}

	return strings.Join([]string{"IssueItemSfV4Domain", string(data)}, " ")
}
