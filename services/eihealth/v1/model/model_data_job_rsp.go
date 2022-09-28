package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据作业列表item
type DataJobRsp struct {

	// 数据作业创建者
	Creator *string `json:"creator,omitempty"`

	// 数据作业结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 数据作业ID
	Id *string `json:"id,omitempty"`

	// 数据作业名称
	Name *string `json:"name,omitempty"`

	// 数据作业创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 数据作业状态
	Status *string `json:"status,omitempty"`

	// 数据作业完成数
	FinishCount *int32 `json:"finish_count,omitempty"`

	// 数据作业总数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 数据作业类型
	Type *string `json:"type,omitempty"`
}

func (o DataJobRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataJobRsp struct{}"
	}

	return strings.Join([]string{"DataJobRsp", string(data)}, " ")
}
