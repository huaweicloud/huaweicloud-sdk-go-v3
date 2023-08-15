package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeVulStatusRequest Request Object
type ChangeVulStatusRequest struct {

	// 企业租户ID，查询所有企业项目时填写：all_granted_eps
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *ChangeVulStatusRequestInfo `json:"body,omitempty"`
}

func (o ChangeVulStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeVulStatusRequest struct{}"
	}

	return strings.Join([]string{"ChangeVulStatusRequest", string(data)}, " ")
}
