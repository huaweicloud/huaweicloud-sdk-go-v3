package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloudStorageLogPathInfo 云存储日志扫描路径信息。
type CloudStorageLogPathInfo struct {

	// 容器挂载路径。
	DirPath *string `json:"dir_path,omitempty"`

	// 日志文件名匹配模式。
	FileNamePattern *string `json:"file_name_pattern,omitempty"`
}

func (o CloudStorageLogPathInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudStorageLogPathInfo struct{}"
	}

	return strings.Join([]string{"CloudStorageLogPathInfo", string(data)}, " ")
}
