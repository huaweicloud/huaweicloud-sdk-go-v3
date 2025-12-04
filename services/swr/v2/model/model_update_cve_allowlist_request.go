package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCveAllowlistRequest 更新漏洞CVE白名单列表的请求
type UpdateCveAllowlistRequest struct {

	// 白名单ID，查询命名空间信息时，会返回白名单ID
	Id *int32 `json:"id,omitempty"`

	// 漏洞白名单列表所属的命名空间ID
	NamespaceId *int32 `json:"namespace_id,omitempty"`

	// 漏洞白名单的有效期时间，以自 1970 年 1 月 1 日以来的秒数表示；如果没有配置，则永不过期
	ExpiresAt *int32 `json:"expires_at,omitempty"`

	// 漏洞列表
	Items *[]CveAllowlistItem `json:"items,omitempty"`
}

func (o UpdateCveAllowlistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCveAllowlistRequest struct{}"
	}

	return strings.Join([]string{"UpdateCveAllowlistRequest", string(data)}, " ")
}
