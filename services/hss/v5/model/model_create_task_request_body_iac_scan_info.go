package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTaskRequestBodyIacScanInfo 创建IAC安全扫描任务的参数，任务类型为“iac_scan”时必传
type CreateTaskRequestBodyIacScanInfo struct {

	// 文件类型，包含如下   - dockerfile：Dockerfile文件   - k8s_yaml：k8s yaml文件
	FileType string `json:"file_type"`

	// 需要扫描的文件id列表，可以从文件上传接口“/v5/{project_id}/common/files/batch-upload”的响应数据中获取上传成功的文件id
	FileIdList []string `json:"file_id_list"`
}

func (o CreateTaskRequestBodyIacScanInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTaskRequestBodyIacScanInfo struct{}"
	}

	return strings.Join([]string{"CreateTaskRequestBodyIacScanInfo", string(data)}, " ")
}
