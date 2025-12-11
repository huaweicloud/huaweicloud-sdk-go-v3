package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsolatedFileRequestInfo 恢复的文件详情
type IsolatedFileRequestInfo struct {

	// **参数解释**： 服务器（主机）的唯一标识ID **取值范围**： 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**： 文件哈希 **取值范围**： 字符长度1-256位
	FileHash *string `json:"file_hash,omitempty"`

	// **参数解释**： 文件路径 **取值范围**： 字符长度1-256位
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释**： 文件的系统属性（如读写权限、隐藏属性、执行权限等） **取值范围**： 字符长度1-256位
	FileAttr *string `json:"file_attr,omitempty"`
}

func (o IsolatedFileRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsolatedFileRequestInfo struct{}"
	}

	return strings.Join([]string{"IsolatedFileRequestInfo", string(data)}, " ")
}
