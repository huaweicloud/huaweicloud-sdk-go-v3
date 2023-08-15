package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrgInstancesRequest Request Object
type ListOrgInstancesRequest struct {

	// 是否页面显示（以标签配置为准）
	IsTemporary *bool `json:"is_temporary,omitempty"`

	// 每页显示的条目数量 10/15/30
	Limit *int64 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询
	Offset *int64 `json:"offset,omitempty"`

	// 租户id（对应华为云帐号的domainId）
	OrgId string `json:"org_id"`

	// 关键字查询(根据实例名，描述模糊查询)
	Search *string `json:"search,omitempty"`
}

func (o ListOrgInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrgInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListOrgInstancesRequest", string(data)}, " ")
}
