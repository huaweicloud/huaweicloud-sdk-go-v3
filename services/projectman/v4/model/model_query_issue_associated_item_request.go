package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryIssueAssociatedItemRequest Request Object
type QueryIssueAssociatedItemRequest struct {

	// 项目32位ID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。
	ProjectId string `json:"project_id"`

	// 工作项唯一ID。可以通过查询工作项列表或者查询树状工作项接口获取，响应消息体中的id字段的值就是工作项ID。
	IssueId string `json:"issue_id"`

	// 工作项类型。
	IssueType string `json:"issue_type"`

	// 项目空间ID，可以通过查询IPD项目列表接口获取，响应消息体中的domain_id字段的值就是项目空间ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 目标项目的32位uuid，项目唯一标识，通过查询IPD项目列表获取，响应消息体中的project_id字段的值就是项目ID。
	TargetProjectId *string `json:"target_project_id,omitempty"`

	// 关联字段的字段编码。
	LinkFieldCode *string `json:"link_field_code,omitempty"`

	// 分页参数，当前页。
	PageNo *string `json:"page_no,omitempty"`

	// 分页参数，页长。
	PageSize *string `json:"page_size,omitempty"`
}

func (o QueryIssueAssociatedItemRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryIssueAssociatedItemRequest struct{}"
	}

	return strings.Join([]string{"QueryIssueAssociatedItemRequest", string(data)}, " ")
}
