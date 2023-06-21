package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 模板报表查询过滤条件。
type ListAimTemplateReportsRequestBody struct {

	// 智能信息模板ID列表，一次最多请求100个。
	TplIds []string `json:"tpl_ids"`

	// 模板报表查询开始时间。样例：2019-10-12T07:20:50.522Z。  >开始时间和结束时间最多间隔90天，超出时间限制返回为空。
	BeginTime string `json:"begin_time"`

	// 模板报表查询结束时间。样例：2019-10-12T07:20:50.522Z。  >开始时间和结束时间最多间隔90天，超出时间限制返回为空。
	EndTime string `json:"end_time"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAimTemplateReportsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimTemplateReportsRequestBody struct{}"
	}

	return strings.Join([]string{"ListAimTemplateReportsRequestBody", string(data)}, " ")
}
