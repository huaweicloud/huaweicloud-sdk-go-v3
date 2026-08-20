package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueAssociateVo 工作项关联请求数据对象
type IssueAssociateVo struct {

	// 关联的工作项ID，多个ID使用逗号分割。可以通过查询工作项列表或者查询树状工作项接口获取，响应消息体中的id字段的值就是工作项ID。 当link_field_code=link时，最多支持关联500个工作项ID，其他场景最多支持50个工作项ID。
	AssociatedIds string `json:"associated_ids"`

	// 操作类型标记位。
	OperationFlag int32 `json:"operation_flag"`

	// 关联项类型编码。
	AssociateIssueType string `json:"associate_issue_type"`

	// 当前工作项类型编码。
	SourceIssueType string `json:"source_issue_type"`

	// 是否使用替换模式。默认为false，追加关联项。如果为true，则会删除原有的关联项，替换为本次关联的工作项。
	IsReplace *bool `json:"is_replace,omitempty"`

	// 关联字段的字段编码。
	LinkFieldCode *string `json:"link_field_code,omitempty"`
}

func (o IssueAssociateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueAssociateVo struct{}"
	}

	return strings.Join([]string{"IssueAssociateVo", string(data)}, " ")
}
