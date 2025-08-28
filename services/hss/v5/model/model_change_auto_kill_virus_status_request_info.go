package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAutoKillVirusStatusRequestInfo 程序自动隔离查杀开关状态请求体
type ChangeAutoKillVirusStatusRequestInfo struct {

	// **参数解释**： 程序自动隔离查杀开关状态 **约束限制**: 不涉及 **取值范围**： - true：开启 - false：关闭  **默认取值**: false
	Enabled *bool `json:"enabled,omitempty"`
}

func (o ChangeAutoKillVirusStatusRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAutoKillVirusStatusRequestInfo struct{}"
	}

	return strings.Join([]string{"ChangeAutoKillVirusStatusRequestInfo", string(data)}, " ")
}
