package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CveAllowlist 漏洞CVE白名单列表
type CveAllowlist struct {

	// 白名单ID
	Id *int32 `json:"id,omitempty"`

	// 漏洞白名单列表所属的命名空间ID
	NamespaceId *int32 `json:"namespace_id,omitempty"`

	// 漏洞白名单的有效期时间，如果没有配置，则永不过期
	ExpiresAt *int32 `json:"expires_at,omitempty"`

	// 漏洞列表
	Items *[]CveAllowlistItem `json:"items,omitempty"`

	// 漏洞白名单的创建时间
	CreationTime *string `json:"creation_time,omitempty"`

	// 漏洞白名单的更新时间
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o CveAllowlist) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CveAllowlist struct{}"
	}

	return strings.Join([]string{"CveAllowlist", string(data)}, " ")
}
