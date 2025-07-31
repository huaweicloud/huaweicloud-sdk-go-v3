package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTasksRequestBodyIacScanInfo IAC安全扫描任务的参数，任务类型task_type为“iac_scan”时可传
type ListTasksRequestBodyIacScanInfo struct {

	// 文件类型，包含如下   - dockerfile：Dockerfile文件   - k8s_yaml：k8s yaml文件
	FileType *string `json:"file_type,omitempty"`
}

func (o ListTasksRequestBodyIacScanInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksRequestBodyIacScanInfo struct{}"
	}

	return strings.Join([]string{"ListTasksRequestBodyIacScanInfo", string(data)}, " ")
}
