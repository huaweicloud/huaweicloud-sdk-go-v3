package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpGroupRequest Request Object
type UpdateIpGroupRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// ip地址组id
	Id string `json:"id"`

	// 增量修改ip地址组时，此为必传字段，传入“add”;删除一个或者多个ip时传入“delete”
	Action *string `json:"action,omitempty"`

	Body *UpdateIpGroupRequestBody `json:"body,omitempty"`
}

func (o UpdateIpGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateIpGroupRequest", string(data)}, " ")
}
