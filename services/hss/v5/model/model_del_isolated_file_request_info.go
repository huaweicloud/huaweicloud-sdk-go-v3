package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DelIsolatedFileRequestInfo 恢复的文件详情
type DelIsolatedFileRequestInfo struct {

	// **参数解释**: 服务器ID **取值范围**: 字符长度1-64位
	HostId string `json:"host_id"`

	// 文件哈希
	FileHash string `json:"file_hash"`

	// 文件路径
	FilePath string `json:"file_path"`

	// 文件属性
	FileAttr string `json:"file_attr"`
}

func (o DelIsolatedFileRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DelIsolatedFileRequestInfo struct{}"
	}

	return strings.Join([]string{"DelIsolatedFileRequestInfo", string(data)}, " ")
}
