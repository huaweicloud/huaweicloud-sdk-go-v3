package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NearLineRecallParam
type NearLineRecallParam struct {

	// 时间过滤。
	TimeLimit *bool `json:"time_limit,omitempty"`

	// 时间特征。
	TimeFeature *string `json:"timeFeature,omitempty"`

	// 保留期(天)。
	RetainDays *int32 `json:"retainDays,omitempty"`

	// 召回字段。
	RecallFileds *[]RecallFiled `json:"recall_fileds,omitempty"`

	// 物品协同过滤作业名称。
	ItemCFJobName *string `json:"itemCF_job_name,omitempty"`
}

func (o NearLineRecallParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NearLineRecallParam struct{}"
	}

	return strings.Join([]string{"NearLineRecallParam", string(data)}, " ")
}
