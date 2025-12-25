package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCodeSegmentResponse Response Object
type CreateCodeSegmentResponse struct {

	// UUID
	CodeSegmentId *string `json:"code_segment_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 代码段名称
	CodeSegmentName *string `json:"code_segment_name,omitempty"`

	// 代码段描述信息
	Description *string `json:"description,omitempty"`

	// 代码片段
	Code *string `json:"code,omitempty"`

	// Iam用户ID
	CreateBy *string `json:"create_by,omitempty"`

	// 毫秒时间戳
	CreateTime *int64 `json:"create_time,omitempty"`

	// Iam用户ID
	UpdateBy *string `json:"update_by,omitempty"`

	// 毫秒时间戳
	UpdateTime     *int64 `json:"update_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateCodeSegmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCodeSegmentResponse struct{}"
	}

	return strings.Join([]string{"CreateCodeSegmentResponse", string(data)}, " ")
}
