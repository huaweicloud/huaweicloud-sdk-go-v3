package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccountsRequest Request Object
type ListAccountsRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// Region ID
	Region string `json:"region"`

	// 每页显示数量
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置
	Offset *int32 `json:"offset,omitempty"`

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListAccountsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccountsRequest struct{}"
	}

	return strings.Join([]string{"ListAccountsRequest", string(data)}, " ")
}
