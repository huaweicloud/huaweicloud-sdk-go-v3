package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteIpdIssueRequest Request Object
type BatchDeleteIpdIssueRequest struct {

	// 项目32位ID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。
	ProjectId string `json:"project_id"`

	// 是否永久删除
	IsPermanentDelete *bool `json:"is_permanent_delete,omitempty"`

	// 工作项的提出项目ID
	SrcProjectId *string `json:"src_project_id,omitempty"`

	Body *[]string `json:"body,omitempty"`
}

func (o BatchDeleteIpdIssueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteIpdIssueRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteIpdIssueRequest", string(data)}, " ")
}
