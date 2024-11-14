package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBackendTargetRequestBody 创建文件系统后端存储请求体
type CreateBackendTargetRequestBody struct {

	// 联动目录名称，SFS Turbo会在文件系统根目录下以该名称创建一个子目录，该目录用于绑定后端存储。  必须是文件系统根目录下不存在的目录名，长度不能超过63个字符，子目录名称不能是“.”或“..”。  不支持多层目录，不能包含字符'/'。
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
