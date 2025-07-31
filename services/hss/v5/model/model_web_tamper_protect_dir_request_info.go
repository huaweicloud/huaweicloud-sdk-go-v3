package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WebTamperProtectDirRequestInfo struct {

	// protect dir list
	ProtectDirList *[]WebTamperProtectHostDirRequestInfo `json:"protect_dir_list,omitempty"`

	// 排除文件类型
	ExcludeFileType *string `json:"exclude_file_type,omitempty"`

	// 排除文件类型
	ExclueFileType *string `json:"exclue_file_type,omitempty"`

	// 防护模式，包含如下4种   - recovery ：恢复模式   - alarm ：告警模式
	ProtectMode *string `json:"protect_mode,omitempty"`
}

func (o WebTamperProtectDirRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebTamperProtectDirRequestInfo struct{}"
	}

	return strings.Join([]string{"WebTamperProtectDirRequestInfo", string(data)}, " ")
}
