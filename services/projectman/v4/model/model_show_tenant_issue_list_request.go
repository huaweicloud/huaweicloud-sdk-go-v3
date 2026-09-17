package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTenantIssueListRequest Request Object
type ShowTenantIssueListRequest struct {

	// 项目32位UUID
	ProjectId *string `json:"project_id,omitempty"`

	// 工作项类型
	IssueType *string `json:"issue_type,omitempty"`

	Body *QueryVo `json:"body,omitempty"`
}

func (o ShowTenantIssueListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantIssueListRequest struct{}"
	}

	return strings.Join([]string{"ShowTenantIssueListRequest", string(data)}, " ")
}
