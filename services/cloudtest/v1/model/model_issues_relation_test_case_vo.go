package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IssuesRelationTestCaseVo struct {

	// 用例ID
	Uri *string `json:"uri,omitempty"`

	// 用例编号
	Number *string `json:"number,omitempty"`

	// 用例名称
	Name *string `json:"name,omitempty"`

	// 工作项id
	DrRelationId *string `json:"dr_relation_id,omitempty"`

	// 状态ID
	StatusCode *string `json:"status_code,omitempty"`

	// 状态名称
	StatusName *string `json:"status_name,omitempty"`

	// 创建人名称
	Author *string `json:"author,omitempty"`

	// 创建人ID
	AuthorId *string `json:"author_id,omitempty"`

	// 处理人名称
	Owner *string `json:"owner,omitempty"`

	// 项目ID
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 创建时间
	CreationDate *string `json:"creation_date,omitempty"`
}

func (o IssuesRelationTestCaseVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssuesRelationTestCaseVo struct{}"
	}

	return strings.Join([]string{"IssuesRelationTestCaseVo", string(data)}, " ")
}
