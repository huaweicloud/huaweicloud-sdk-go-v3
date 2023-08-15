package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStudyJobReq 创建study作业请求体
type CreateStudyJobReq struct {

	// workflow作业id
	WorkflowJobId string `json:"workflow_job_id"`

	// 数据库模板id
	TemplateId *string `json:"template_id,omitempty"`

	// 数据库名称
	DatabaseName string `json:"database_name"`

	// 生成数据库实例的文件相对路径
	RelativePath string `json:"relative_path"`

	OutputFileType *OutputFileType `json:"output_file_type"`
}

func (o CreateStudyJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStudyJobReq struct{}"
	}

	return strings.Join([]string{"CreateStudyJobReq", string(data)}, " ")
}
