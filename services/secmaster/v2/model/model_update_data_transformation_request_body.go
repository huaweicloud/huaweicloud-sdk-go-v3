package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDataTransformationRequestBody struct {

	// 数据转换名称
	DataTransformationName *string `json:"data_transformation_name,omitempty"`

	// 数据转换描述
	Description *string `json:"description,omitempty"`

	// directory 目录分组
	Directory *string `json:"directory,omitempty"`

	// Job Script 作业脚本
	Script *string `json:"script,omitempty"`

	Status *JobStatus `json:"status,omitempty"`

	// 归属
	Belong *string `json:"belong,omitempty"`

	JobMode *IsapJobMode `json:"job_mode,omitempty"`

	// cu总量
	CuQuotaAmount *float32 `json:"cu_quota_amount,omitempty"`

	JobModeSetting *IsapJobModeSettingDto `json:"job_mode_setting,omitempty"`

	Environment *JobEnvironment `json:"environment,omitempty"`

	// UUID
	OutputTableId *string `json:"output_table_id,omitempty"`

	// 输出表ID列表
	OutputTableIds *[]string `json:"output_table_ids,omitempty"`

	// 输出表名称列表
	OutputTableNames *[]string `json:"output_table_names,omitempty"`
}

func (o UpdateDataTransformationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataTransformationRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDataTransformationRequestBody", string(data)}, " ")
}
