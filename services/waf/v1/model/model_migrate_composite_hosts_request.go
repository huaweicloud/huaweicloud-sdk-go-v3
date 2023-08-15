package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrateCompositeHostsRequest Request Object
type MigrateCompositeHostsRequest struct {

	// 当前企业项目ID
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 迁移的目标企业项目ID
	TargetEnterpriseProjectId string `json:"target_enterprise_project_id"`

	Body *MigrateCompositeHostsRequestBody `json:"body,omitempty"`
}

func (o MigrateCompositeHostsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateCompositeHostsRequest struct{}"
	}

	return strings.Join([]string{"MigrateCompositeHostsRequest", string(data)}, " ")
}
