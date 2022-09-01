package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type NearLineRecallParam struct {

	// 时间过滤。
	TimeLimit *bool `json:"time_limit,omitempty" xml:"time_limit"`

	// 时间特征。
	TimeFeature *string `json:"timeFeature,omitempty" xml:"timeFeature"`

	// 保留期(天)。
	RetainDays *int32 `json:"retainDays,omitempty" xml:"retainDays"`

	// 召回字段。
	RecallFileds *[]RecallFiled `json:"recall_fileds,omitempty" xml:"recall_fileds"`

	// 物品协同过滤作业名称。
	ItemCFJobName *string `json:"itemCF_job_name,omitempty" xml:"itemCF_job_name"`
}

func (o NearLineRecallParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NearLineRecallParam struct{}"
	}

	return strings.Join([]string{"NearLineRecallParam", string(data)}, " ")
}
