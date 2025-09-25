package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WtpProtectDirResponseInfo 防护目录信息
type WtpProtectDirResponseInfo struct {

	// **参数解释**: 防护目录列表 **取值范围**: 最少0条，最多50条
	ProtectDirList *[]WtpProtectHostDirResponseInfo `json:"protect_dir_list,omitempty"`

	// **参数解释**: 排除文件类型 **取值范围**: 字符长度0-512位
	ExcludeFileType *string `json:"exclude_file_type,omitempty"`

	// **参数解释**: 防护模式。 **取值范围**: - recovery ：拦截模式。 - alarm ：告警模式，仅Linux服务器支持告警模式。
	ProtectMode *string `json:"protect_mode,omitempty"`
}

func (o WtpProtectDirResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WtpProtectDirResponseInfo struct{}"
	}

	return strings.Join([]string{"WtpProtectDirResponseInfo", string(data)}, " ")
}
