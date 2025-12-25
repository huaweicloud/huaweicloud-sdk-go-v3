package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDataTransformationRequestBody struct {

	// 数据转换名称
	DataTransformationName string `json:"data_transformation_name"`

	// 数据转换描述
	Description string `json:"description"`

	// directory 目录分组
	Directory *string `json:"directory,omitempty"`

	// Job Script 作业脚本
	Script string `json:"script"`

	Status *JobStatus `json:"status"`

	JobMode *IsapJobMode `json:"job_mode"`

	JobModeSetting *IsapJobModeSettingDto `json:"job_mode_setting,omitempty"`

	Environment *JobEnvironment `json:"environment"`

	// UUID
	OutputTableId *string `json:"output_table_id,omitempty"`

	// 归属
	Belong *string `json:"belong,omitempty"`

	// 数量
	CuQuotaAmount float32 `json:"cu_quota_amount"`

	// 输入表ID
	InputTableId *[]string `json:"input_table_id,omitempty"`

	// 输入表名称列表
	InputTableNames *[]string `json:"input_table_names,omitempty"`

	// 输出表ID列表
	OutputTableIds []string `json:"output_table_ids"`

	// 输出表名称列表
	OutputTableNames []string `json:"output_table_names"`
}

func (o CreateDataTransformationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataTransformationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDataTransformationRequestBody", string(data)}, " ")
}
