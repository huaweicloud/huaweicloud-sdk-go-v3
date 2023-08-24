package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResolveTasksRequest Request Object
type ListResolveTasksRequest struct {

	// 解析任务ID。
	TaskId *string `json:"task_id,omitempty"`

	// 任务名称。
	TaskName *string `json:"task_name,omitempty"`

	// 智能信息模板ID。
	TplId *string `json:"tpl_id,omitempty"`

	// 智能信息模板名称。
	TplName *string `json:"tpl_name,omitempty"`

	// 创建解析任务时填写用户唯一标识，手机号码或者任何的唯一标识，唯一标识不超过64个字符。 发送智能信息时则必须填客户的手机号码。样例为：130****0001。
	CustFlag *string `json:"cust_flag,omitempty"`

	// 完整的短链连接地址。样例：km2g.cn/PDiWqc。
	AimUrl *string `json:"aim_url,omitempty"`

	//  智能信息解析任务创建开始时间。格式为：2019-10-12T07:20:50.522Z。  > 需同时传入end_time才能生效，单独传begin_time不会作为过滤条件。 > > 若不填，则默认查询24小时内创建的解析任务。
	BeginTime *string `json:"begin_time,omitempty"`

	//  智能信息解析任务创建结束时间。格式为：2019-10-12T07:20:50.522Z。  > 需同时传入begin_time才能生效，单独传end_time不会作为过滤条件。 > > 若不填，则默认查询24小时内创建的解析任务。
	EndTime *string `json:"end_time,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0。  >为提高查询效率，offset+limit须小于等于10000，超出范围查询为空。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListResolveTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResolveTasksRequest struct{}"
	}

	return strings.Join([]string{"ListResolveTasksRequest", string(data)}, " ")
}
