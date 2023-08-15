package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataJobRequest Request Object
type ListDataJobRequest struct {

	// 创建者名称
	Creator *string `json:"creator,omitempty"`

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 查询该时间之后创建的数据作业
	FromTime *int64 `json:"from_time,omitempty"`

	// 查询条数
	Limit *int32 `json:"limit,omitempty"`

	// 数据作业名称
	Name *string `json:"name,omitempty"`

	// 查询偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 数据作业状态
	Status *string `json:"status,omitempty"`

	// 查询该时间之前创建的数据作业
	ToTime *int64 `json:"to_time,omitempty"`

	// 数据作业类型
	Type *string `json:"type,omitempty"`

	// 查询该时间之后完成的数据作业
	FinishFromTime *int64 `json:"finish_from_time,omitempty"`

	// 查询该时间之前完成的数据作业
	FinishToTime *int64 `json:"finish_to_time,omitempty"`

	// 排序规则 目前默认时间降序
	SortDir *string `json:"sort_dir,omitempty"`

	// 排序规则 目前默认时间降序，支持根据status,name,type,creator,create_time,end_time
	SortKey *string `json:"sort_key,omitempty"`
}

func (o ListDataJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataJobRequest struct{}"
	}

	return strings.Join([]string{"ListDataJobRequest", string(data)}, " ")
}
