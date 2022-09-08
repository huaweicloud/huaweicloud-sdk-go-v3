package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateBatchJobResponse struct {

	// 仅在响应返回。作业ID。
	JobId *string `json:"job_id,omitempty"`

	// 作业名称。只能包含数字、英文字母、中文字母、下划线以及中划线。长度为1~128。
	JobName *string `json:"job_name,omitempty"`

	// 标签。只能包含数字、英文字母、中文字符、下划线、中划线、逗号以及斜杠。长度为0~128。
	Tags *string `json:"tags,omitempty"`

	// 仅在响应返回。创建时间。
	CreatedTime *string `json:"created_time,omitempty"`

	// 仅在响应返回。更新时间。
	ModifiedTime *string `json:"modified_time,omitempty"`

	// 作业类型。目前仅支持SqlJob.
	JobType *string `json:"job_type,omitempty"`

	// 作业查询结果导出到OBS的路径。覆写已存在文件。
	ExportPath *string `json:"export_path,omitempty"`

	// 导出文件时是否合并结果文件。true：合并成一个结果文件；false：不合并结果文件。
	MergeResultFile *bool `json:"merge_result_file,omitempty"`

	SqlJob *SqlJob `json:"sql_job,omitempty"`

	Schedule       *Schedule `json:"schedule,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateBatchJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBatchJobResponse struct{}"
	}

	return strings.Join([]string{"CreateBatchJobResponse", string(data)}, " ")
}
