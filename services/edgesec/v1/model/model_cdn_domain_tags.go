package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CdnDomainTags cdn域名安全服务约束信息
type CdnDomainTags struct {

	// 约束原因
	Notes *string `json:"notes,omitempty"`

	// 约束内容
	Constraint *string `json:"constraint,omitempty"`
}

func (o CdnDomainTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdnDomainTags struct{}"
	}

	return strings.Join([]string{"CdnDomainTags", string(data)}, " ")
}
