package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkspacesRequest Request Object
type ListWorkspacesRequest struct {

	// 偏移量，表示从此偏移量开始查询， offset大于等于0，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 指定每一页返回的最大条目数，取值范围[1,100]，默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// 使用名称进行检索，支持模糊查询。只能包含字母、数字、下划线和中划线，且长度为4到32个字符。
	Name *string `json:"name,omitempty"`

	// 工作空间Id
	Id *string `json:"id,omitempty"`

	// 企业项目ID。由用户输入，只能包含字母、数字、下划线和中划线，且长度为1到36个字符。默认为all_granted_eps
	EnterpriseProjectId string `json:"enterprise_project_id"`
}

func (o ListWorkspacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkspacesRequest struct{}"
	}

	return strings.Join([]string{"ListWorkspacesRequest", string(data)}, " ")
}
