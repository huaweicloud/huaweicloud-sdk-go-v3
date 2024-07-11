package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnvironmentHostsRequest Request Object
type ListEnvironmentHostsRequest struct {

	// 应用id
	ApplicationId string `json:"application_id"`

	// 环境id
	EnvironmentId string `json:"environment_id"`

	// 主机名、ip关键字模糊搜索
	KeyField *string `json:"key_field,omitempty"`

	// 是否为代理机,true为代理机
	AsProxy *bool `json:"as_proxy,omitempty"`

	// 分页页码
	PageIndex *int32 `json:"page_index,omitempty"`

	// 分页查询每页条数
	PageSize *int32 `json:"page_size,omitempty"`
}

func (o ListEnvironmentHostsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvironmentHostsRequest struct{}"
	}

	return strings.Join([]string{"ListEnvironmentHostsRequest", string(data)}, " ")
}
