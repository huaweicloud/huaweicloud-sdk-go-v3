package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AutoJobListDto struct {

	// 自动作业id
	Id *string `json:"id,omitempty"`

	// 自动作业的名称，取值范围：[1,63]，允许大小写字母、数字、以及特殊字符中划线(-)
	Name *string `json:"name,omitempty"`

	// 自动作业的描述, 取值范围：输入字符最大长度为255
	Description *string `json:"description,omitempty"`

	// 自动作业标签
	Labels *[]string `json:"labels,omitempty"`

	// 自动作业优先级，[0,9]，0表示最低，默认0
	Priority *int32 `json:"priority,omitempty"`

	// 作业执行超时时长，取值范围: [1, 144000]，单位：分钟，默认数值1440
	Timeout *int32 `json:"timeout,omitempty"`

	// 自动作业状态
	Status *string `json:"status,omitempty"`

	// 自动作业创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 自动作业结束时间
	FinishTime *string `json:"finish_time,omitempty"`

	// 失败原因，当自动作业执行失败时会返回，比如依赖的数据表，流程不存在等等
	FailedReason *string `json:"failed_reason,omitempty"`

	// 自动作业的创建者
	UserName *string `json:"user_name,omitempty"`

	ToolInfo *ToolInfoDto `json:"tool_info,omitempty"`

	// 自动作业依赖的数据表ID
	DatabaseId *string `json:"database_id,omitempty"`

	// 自动作业依赖的数据表名称
	DatabaseName *string `json:"database_name,omitempty"`
}

func (o AutoJobListDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoJobListDto struct{}"
	}

	return strings.Join([]string{"AutoJobListDto", string(data)}, " ")
}
