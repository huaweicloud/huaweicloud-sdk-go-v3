package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBackendTargetRequestBody 创建文件系统后端存储库请求体
type CreateBackendTargetRequestBody struct {

	// 文件系统路径
	FileSystemPath string `json:"file_system_path"`

	Obs *ObsDataRepository `json:"obs"`
}

func (o CreateBackendTargetRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBackendTargetRequestBody struct{}"
	}

	return strings.Join([]string{"CreateBackendTargetRequestBody", string(data)}, " ")
}
