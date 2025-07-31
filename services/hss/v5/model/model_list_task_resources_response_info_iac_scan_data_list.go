package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskResourcesResponseInfoIacScanDataList 任务下的文件扫描详情信息
type ListTaskResourcesResponseInfoIacScanDataList struct {

	// 文件名称
	FileName *string `json:"file_name,omitempty"`

	// 文件类型，包含如下   - dockerfile：Dockerfile文件   - k8s_yaml：k8s yaml文件
	FileType *string `json:"file_type,omitempty"`

	// chart包名称
	ChartName *string `json:"chart_name,omitempty"`

	// 文件ID，服务生成的文件唯一ID
	FileId *string `json:"file_id,omitempty"`

	// 文件大小
	FileSize *int32 `json:"file_size,omitempty"`

	// 文件的扫描状态，包含如下：   - scanning：扫描中   - success：扫描成功   - failed：扫描失败
	ScanStatus *string `json:"scan_status,omitempty"`

	// 扫描失败的原因
	FailedReason *string `json:"failed_reason,omitempty"`

	// 扫描开始的时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 扫描结束的时间
	EndTime *int64 `json:"end_time,omitempty"`
}

func (o ListTaskResourcesResponseInfoIacScanDataList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskResourcesResponseInfoIacScanDataList struct{}"
	}

	return strings.Join([]string{"ListTaskResourcesResponseInfoIacScanDataList", string(data)}, " ")
}
