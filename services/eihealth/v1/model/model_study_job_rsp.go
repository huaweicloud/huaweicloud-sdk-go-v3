package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StudyJobRsp 作业详情返回体
type StudyJobRsp struct {

	// study作业id
	Id *string `json:"id,omitempty"`

	// workflow作业id
	WorkflowJobId *string `json:"workflow_job_id,omitempty"`

	// 作业名称
	Name *string `json:"name,omitempty"`

	// 作业状态
	Status *string `json:"status,omitempty"`

	// 生成study作业结果的模板id
	TemplateId *string `json:"template_id,omitempty"`

	// study作业结果的数据库实例名称
	DatabaseName *string `json:"database_name,omitempty"`

	// study作业结果的数据库实例id
	DatabaseId *string `json:"database_id,omitempty"`

	// 生成study作业结果的文件的相对路径
	RelativePath *string `json:"relative_path,omitempty"`

	// 生成study作业结果的文件的类型
	OutputFileType *string `json:"output_file_type,omitempty"`

	// 使用的workflow名称
	WorkflowName *string `json:"workflow_name,omitempty"`

	// 使用的workflow标签
	Label *string `json:"label,omitempty"`

	// 作业创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 作业更新时间
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o StudyJobRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StudyJobRsp struct{}"
	}

	return strings.Join([]string{"StudyJobRsp", string(data)}, " ")
}
