package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTasksResponseInfoIacScanInfo IAC安全扫描任务特有的任务信息
type ListTasksResponseInfoIacScanInfo struct {

	// 文件类型，包含如下   - dockerfile：Dockerfile文件   - k8s_yaml：k8s yaml文件
	FileType *string `json:"file_type,omitempty"`

	// 当前任务扫描的文件总数
	ScanFileNum *int32 `json:"scan_file_num,omitempty"`

	// 当前任务扫描成功的文件数量
	SuccessFileNum *int32 `json:"success_file_num,omitempty"`

	// 当前任务扫描失败的文件数量
	FailedFileNum *int32 `json:"failed_file_num,omitempty"`
}

func (o ListTasksResponseInfoIacScanInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksResponseInfoIacScanInfo struct{}"
	}

	return strings.Join([]string{"ListTasksResponseInfoIacScanInfo", string(data)}, " ")
}
