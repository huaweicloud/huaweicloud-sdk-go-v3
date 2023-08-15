package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeHostsGroupRequest Request Object
type ChangeHostsGroupRequest struct {

	// region id
	Region string `json:"region"`

	// 企业项目ID，查询所有企业项目时填写：all_granted_eps
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *ChangeHostsGroupRequestInfo `json:"body,omitempty"`
}

func (o ChangeHostsGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeHostsGroupRequest struct{}"
	}

	return strings.Join([]string{"ChangeHostsGroupRequest", string(data)}, " ")
}
