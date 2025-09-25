package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebTamperProtectDirRequestInfo 防护目录信息
type WebTamperProtectDirRequestInfo struct {

	// **参数解释**: 防护目录列表 **约束限制**: 不涉及 **取值范围**: 最少1条，最多50条 **默认取值**: 不涉及
	ProtectDirList []WebTamperProtectHostDirRequestInfo `json:"protect_dir_list"`

	// **参数解释**: 排除文件类型 **约束限制**: 不涉及 **取值范围**: 文件类型只能输入字母、数字，最多支持10个文件类型，每个文件类型长度不超过10个字符，多个文件类型以分号隔开。 **默认取值**: 不涉及
	ExcludeFileType *string `json:"exclude_file_type,omitempty"`

	// **参数解释**: 防护模式，仅Linux服务器支持设置防护模式为告警模式，Windows服务器仅支持拦截模式。 **约束限制**: 不涉及 **取值范围**: - recovery ：拦截模式。 - alarm ：告警模式，仅Linux服务器支持。  **默认取值**: recovery
	ProtectMode *string `json:"protect_mode,omitempty"`
}

func (o WebTamperProtectDirRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebTamperProtectDirRequestInfo struct{}"
	}

	return strings.Join([]string{"WebTamperProtectDirRequestInfo", string(data)}, " ")
}
