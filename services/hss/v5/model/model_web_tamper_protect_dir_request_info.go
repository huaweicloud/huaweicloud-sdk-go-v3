package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WebTamperProtectDirRequestInfo struct {

	// 防护目录列表
	ProtectDirList *[]WebTamperProtectHostDirRequestInfo `json:"protect_dir_list,omitempty"`

	// 排除文件类型
	ExcludeFileType *string `json:"exclude_file_type,omitempty"`

	// 排除文件类型
	ExclueFileType *string `json:"exclue_file_type,omitempty"`

	// **参数解释**: 防护模式 **约束限制**: 不涉及 **取值范围**: - recovery ：拦截模式。 - alarm ：告警模式。  **默认取值**: recovery
	ProtectMode *string `json:"protect_mode,omitempty"`
}

func (o WebTamperProtectDirRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebTamperProtectDirRequestInfo struct{}"
	}

	return strings.Join([]string{"WebTamperProtectDirRequestInfo", string(data)}, " ")
}
