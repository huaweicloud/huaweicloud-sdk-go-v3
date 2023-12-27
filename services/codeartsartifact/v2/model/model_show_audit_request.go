package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuditRequest Request Object
type ShowAuditRequest struct {

	// 租户id
	TenantId string `json:"tenant_id"`

	// 项目id
	ProjectId string `json:"project_id"`

	// 模块
	Module string `json:"module"`

	// 仓库id
	Repo string `json:"repo"`

	// 用户id
	UserId *string `json:"user_id,omitempty"`

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 页码
	PageNum *int32 `json:"page_num,omitempty"`

	// 每页大小
	PageSize *int32 `json:"page_size,omitempty"`
}

func (o ShowAuditRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditRequest struct{}"
	}

	return strings.Join([]string{"ShowAuditRequest", string(data)}, " ")
}
