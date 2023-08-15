package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddHostsGroupRequest Request Object
type AddHostsGroupRequest struct {

	// region id
	Region string `json:"region"`

	// 企业项目ID，查询所有企业项目时填写：all_granted_eps
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *AddHostsGroupRequestInfo `json:"body,omitempty"`
}

func (o AddHostsGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddHostsGroupRequest struct{}"
	}

	return strings.Join([]string{"AddHostsGroupRequest", string(data)}, " ")
}
