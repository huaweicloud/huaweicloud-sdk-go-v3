package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetProtectDirSwitchRequestInfo struct {

	// **参数解释**: 需要暂停或恢复的防护目录列表 **约束限制**: 不涉及 **取值范围**: 不超过50条 **默认取值**: 不涉及
	ProtectDirList []string `json:"protect_dir_list"`

	// **参数解释**: 暂停或恢复防护目录的防护状态 **约束限制**: 不涉及 **取值范围**: - true ：恢复防护。 - false ：暂停防护。  **默认取值**: false
	EnableProtect bool `json:"enable_protect"`
}

func (o SetProtectDirSwitchRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetProtectDirSwitchRequestInfo struct{}"
	}

	return strings.Join([]string{"SetProtectDirSwitchRequestInfo", string(data)}, " ")
}
