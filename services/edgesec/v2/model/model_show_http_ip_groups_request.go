package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpIpGroupsRequest Request Object
type ShowHttpIpGroupsRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id，默认为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 分页查询参数，第page页
	Page *int32 `json:"page,omitempty"`

	// 分页查询参数，每页显示pagesize条记录
	Pagesize *int32 `json:"pagesize,omitempty"`

	// IP地址组名称
	Name *string `json:"name,omitempty"`

	// IP地址/地址段
	Ip *string `json:"ip,omitempty"`
}

func (o ShowHttpIpGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpIpGroupsRequest struct{}"
	}

	return strings.Join([]string{"ShowHttpIpGroupsRequest", string(data)}, " ")
}
