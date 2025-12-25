package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCloudLogTenantWhitelistRequest Request Object
type ShowCloudLogTenantWhitelistRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 账号ID
	DomainId string `json:"domain_id"`

	// 第几页
	Offset *int64 `json:"offset,omitempty"`

	// 每页数量
	Limit *int64 `json:"limit,omitempty"`

	// 排序顺序
	SortDir *string `json:"sort_dir,omitempty"`

	// 排序字段
	SortKey *string `json:"sort_key,omitempty"`
}

func (o ShowCloudLogTenantWhitelistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCloudLogTenantWhitelistRequest struct{}"
	}

	return strings.Join([]string{"ShowCloudLogTenantWhitelistRequest", string(data)}, " ")
}
