package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListExecutionResponseData struct {

	// id
	Id *string `json:"id,omitempty"`

	// 工单唯一id
	ExecutionId *string `json:"execution_id,omitempty"`

	// 作业名称
	DocumentName *string `json:"document_name,omitempty"`

	// 作业id
	DocumentId *string `json:"document_id,omitempty"`

	// 作业版本id
	DocumentVersionId *string `json:"document_version_id,omitempty"`

	// 作业版本号
	DocumentVersion *string `json:"document_version,omitempty"`

	// 工单执行开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 工单执行结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 工单更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 工单创建时间
	Creator *string `json:"creator,omitempty"`

	// 工单状态
	Status *string `json:"status,omitempty"`

	// 工单执行描述
	Description *string `json:"description,omitempty"`

	// 工单执行全局参数
	Parameters *[]ListExecutionResponseParameters `json:"parameters,omitempty"`

	// 系统标签列表
	SysTags *[]CreateDocumentRequestBodyTags `json:"sys_tags,omitempty"`

	// 自定义标签列表
	Tags *[]CreateDocumentRequestBodyTags `json:"tags,omitempty"`

	// 工单类型：BATCH、RATE_CONTROL、NORMAL
	Type *string `json:"type,omitempty"`

	// 速率模式执行指定参数
	TargetParameterName *string `json:"target_parameter_name,omitempty"`

	// 速率模式执行指定元素
	Targets *[]Target `json:"targets,omitempty"`
}

func (o ListExecutionResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExecutionResponseData struct{}"
	}

	return strings.Join([]string{"ListExecutionResponseData", string(data)}, " ")
}
