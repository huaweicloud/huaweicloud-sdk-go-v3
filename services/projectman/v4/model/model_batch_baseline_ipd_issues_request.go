package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchBaselineIpdIssuesRequest Request Object
type BatchBaselineIpdIssuesRequest struct {

	// 项目32位ID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。
	ProjectId string `json:"project_id"`

	Body *BatchBaselineIpdIssuesParam `json:"body,omitempty"`
}

func (o BatchBaselineIpdIssuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchBaselineIpdIssuesRequest struct{}"
	}

	return strings.Join([]string{"BatchBaselineIpdIssuesRequest", string(data)}, " ")
}
