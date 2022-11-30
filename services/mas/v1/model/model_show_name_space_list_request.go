package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowNameSpaceListRequest struct {

	// 偏移量
	Offset *string `json:"offset,omitempty"`

	// 每页显示的条目数量
	Limit *string `json:"limit,omitempty"`

	// 命名空间名称
	Name *string `json:"name,omitempty"`

	// 多活类型，1：同城多活，2：异地多活
	Type *string `json:"type,omitempty"`

	// 是否已被使用   true  ：是   false  ：否
	IsUsed *string `json:"is_used,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowNameSpaceListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNameSpaceListRequest struct{}"
	}

	return strings.Join([]string{"ShowNameSpaceListRequest", string(data)}, " ")
}
