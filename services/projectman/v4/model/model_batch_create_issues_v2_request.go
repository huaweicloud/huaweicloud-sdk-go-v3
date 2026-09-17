package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateIssuesV2Request Request Object
type BatchCreateIssuesV2Request struct {

	// 项目32位ID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。
	ProjectId string `json:"project_id"`

	Body *[]IssueCreateEntity `json:"body,omitempty"`
}

func (o BatchCreateIssuesV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateIssuesV2Request struct{}"
	}

	return strings.Join([]string{"BatchCreateIssuesV2Request", string(data)}, " ")
}
