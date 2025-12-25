package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDataTransformationResponse Response Object
type CreateDataTransformationResponse struct {

	// UUID
	DataTransformationId *string `json:"data_transformation_id,omitempty"`

	// 数据转换名称
	DataTransformationName *string `json:"data_transformation_name,omitempty"`

	// Job Script 作业脚本
	Script *string `json:"script,omitempty"`

	Status *JobStatus `json:"status,omitempty"`

	// directory 目录分组
	Directory *string `json:"directory,omitempty"`

	// 数据转换描述
	Description *string `json:"description,omitempty"`

	JobMode *IsapJobMode `json:"job_mode,omitempty"`

	JobModeSetting *IsapJobModeSettingVo `json:"job_mode_setting,omitempty"`

	ProcessStatus *JobProcessStatus `json:"process_status,omitempty"`

	ProcessError *DataTransformationProcessError `json:"process_error,omitempty"`

	Environment *JobEnvironment `json:"environment,omitempty"`

	// UUID
	OutputTableId *string `json:"output_table_id,omitempty"`

	// 表名称
	OutputTableName *string `json:"output_table_name,omitempty"`

	// 输出表ID列表
	OutputTableIds *[]string `json:"output_table_ids,omitempty"`

	// 输出表名称列表
	OutputTableNames *[]string `json:"output_table_names,omitempty"`

	// 创建者
	CreateBy *string `json:"create_by,omitempty"`

	// 毫秒时间戳
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新者
	UpdateBy *string `json:"update_by,omitempty"`

	// 毫秒时间戳
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 毫秒时间戳
	DeleteTime     *int64 `json:"delete_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateDataTransformationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataTransformationResponse struct{}"
	}

	return strings.Join([]string{"CreateDataTransformationResponse", string(data)}, " ")
}
