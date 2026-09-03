package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTasksBySqlTemplateIdRequestBody 按SQL模板ID查询全量SQL任务请求体
type ListTasksBySqlTemplateIdRequestBody struct {

	// 实例ID，实例的唯一标识
	InstanceId string `json:"instance_id"`

	// 节点ID，实例节点的唯一标识
	NodeId *string `json:"node_id,omitempty"`

	// 起止时间的查询左区间
	RangeLeft *int64 `json:"range_left,omitempty"`

	// 起止时间的查询右区间
	RangeRight *int64 `json:"range_right,omitempty"`

	// SQL模板ID
	SqlTemplateId string `json:"sql_template_id"`

	// 每页记录数
	PageSize *int32 `json:"page_size,omitempty"`

	// 当前页码
	CurPage *int32 `json:"cur_page,omitempty"`
}

func (o ListTasksBySqlTemplateIdRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksBySqlTemplateIdRequestBody struct{}"
	}

	return strings.Join([]string{"ListTasksBySqlTemplateIdRequestBody", string(data)}, " ")
}
