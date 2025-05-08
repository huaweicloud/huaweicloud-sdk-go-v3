package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckScriptRiskResDataBlacklistCommands struct {

	// 命令。
	CommandName *string `json:"command_name,omitempty"`

	// 匹配规则。
	CommandRule *string `json:"command_rule,omitempty"`

	// 实例。
	Example *string `json:"example,omitempty"`

	// 黑名单命令中文描述。
	DescriptionEn *string `json:"description_en,omitempty"`

	// 黑名单命令英文描述。
	DescriptionZh *string `json:"description_zh,omitempty"`
}

func (o CheckScriptRiskResDataBlacklistCommands) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckScriptRiskResDataBlacklistCommands struct{}"
	}

	return strings.Join([]string{"CheckScriptRiskResDataBlacklistCommands", string(data)}, " ")
}
