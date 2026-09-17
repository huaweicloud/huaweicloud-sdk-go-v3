package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateIssuesParam 批量更新工作项参数，支持更新不同工作项的相同字段
type BatchUpdateIssuesParam struct {

	// **参数解释**： 需要更新的工作项ID数组，可通过[高级查询工作项](ListIssuesV4.xml)接口获取，响应消息体中的**id**字段的值就是工作项ID。 **约束限制**： 18~19位的数字字符串(工作项的**id**字段对应的字符串)。
	Id []string `json:"id"`

	Attribute *IssueUpdateAttribute `json:"attribute"`
}

func (o BatchUpdateIssuesParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateIssuesParam struct{}"
	}

	return strings.Join([]string{"BatchUpdateIssuesParam", string(data)}, " ")
}
