package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReplayReportExportStatusResponse Response Object
type ShowReplayReportExportStatusResponse struct {

	// 导出状态。取值范围： - EXPORTING ：导出中 - EXPORT_COMPLETE ：导出完成 - EXPORT_COMMON_FAILED ：导出失败
	ExportStatus *string `json:"export_status,omitempty"`

	// 任务id
	JobId *string `json:"job_id,omitempty"`

	// 导出的sql文件类型。取值范围： - abnormal_sql ：异常sql列表 - error_sql_detail ：异常sql详情 - slow_sql ：慢sql列表 - slow_sql_detail ： 慢sql详情
	FileType *string `json:"file_type,omitempty"`

	// 失败原因
	FailedReason *string `json:"failed_reason,omitempty"`

	// 导出的数据总量
	TotalCount *int64 `json:"total_count,omitempty"`

	// 当前已经处理数据量
	CurrentCount *int64 `json:"current_count,omitempty"`

	// 任务进度百分数
	ProgressPercentage *int32 `json:"progress_percentage,omitempty"`

	// 已经上传到obs的文件名称
	UploadedFileNames *[]string `json:"uploaded_file_names,omitempty"`
	HttpStatusCode    int       `json:"-"`
}

func (o ShowReplayReportExportStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplayReportExportStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowReplayReportExportStatusResponse", string(data)}, " ")
}
