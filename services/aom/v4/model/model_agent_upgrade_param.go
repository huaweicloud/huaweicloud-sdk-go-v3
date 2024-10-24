package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgentUpgradeParam struct {

	// UniAgent升级的版本号。
	Version string `json:"version"`

	// UniAgent主机列表信息。
	AgentList []SingleAgentParam `json:"agent_list"`
}

func (o AgentUpgradeParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentUpgradeParam struct{}"
	}

	return strings.Join([]string{"AgentUpgradeParam", string(data)}, " ")
}
