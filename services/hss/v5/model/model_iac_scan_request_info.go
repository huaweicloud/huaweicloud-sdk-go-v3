package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IacScanRequestInfo iac扫描信息
type IacScanRequestInfo struct {

	// **参数解释**: 配置方式 **约束限制**: 不涉及 **取值范围**: - local_directory：本地目录。 - remote_address：https远程地址。 - git_repository_address：git仓库地址。  **默认取值**: 不涉及
	ConfigurationMode *string `json:"configuration_mode,omitempty"`

	// **参数解释**: 配置文件路径 **约束限制**: 不涉及 **取值范围**: 字符取值1-256 **默认取值**: 不涉及
	Path *string `json:"path,omitempty"`
}

func (o IacScanRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IacScanRequestInfo struct{}"
	}

	return strings.Join([]string{"IacScanRequestInfo", string(data)}, " ")
}
