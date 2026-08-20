package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateIssuesRequest Request Object
type AssociateIssuesRequest struct {

	// 项目32位ID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。
	ProjectId string `json:"project_id"`

	// 工作项唯一ID。可以通过查询工作项列表或者查询树状工作项接口获取，响应消息体中的id字段的值就是工作项ID。
	IssueId string `json:"issue_id"`

	// 目标项目的32位UUID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。非跨项目场景，与请求路径中的project_id一致。
	DstDomainId string `json:"dst_domain_id"`

	Body *IssueAssociateVo `json:"body,omitempty"`
}

func (o AssociateIssuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateIssuesRequest struct{}"
	}

	return strings.Join([]string{"AssociateIssuesRequest", string(data)}, " ")
}
