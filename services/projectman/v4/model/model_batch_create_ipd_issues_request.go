package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateIpdIssuesRequest Request Object
type BatchCreateIpdIssuesRequest struct {

	// 项目32位ID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。
	ProjectId string `json:"project_id"`

	Body *[]CreateIpdProjectIssueParam `json:"body,omitempty"`
}

func (o BatchCreateIpdIssuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateIpdIssuesRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateIpdIssuesRequest", string(data)}, " ")
}
