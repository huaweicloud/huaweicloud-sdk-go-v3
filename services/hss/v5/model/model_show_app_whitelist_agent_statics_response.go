package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppWhitelistAgentStaticsResponse Response Object
type ShowAppWhitelistAgentStaticsResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	AgentNum       *int32 `json:"agent_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAppWhitelistAgentStaticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppWhitelistAgentStaticsResponse struct{}"
	}

	return strings.Join([]string{"ShowAppWhitelistAgentStaticsResponse", string(data)}, " ")
}
