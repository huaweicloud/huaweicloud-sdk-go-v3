package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CloseProtectionInfoRequestInfo struct {

	// **参数解释**: 需要关闭勒索防护的主机ID列表 **约束限制**: 不涉及 **取值范围**: 列表条数0-64 **默认取值**: 不涉及
	HostIdList []string `json:"host_id_list"`

	// **参数解释**: 需要关闭勒索防护的agentID列表 **约束限制**: 不涉及 **取值范围**: 列表条数0-64 **默认取值**: 不涉及
	AgentIdList []string `json:"agent_id_list"`

	// **参数解释**: 关闭防护类型，包含如下：   - close_anti : 关闭勒索防护；暂不支持关闭备份防护，若需要解绑存储库，请前往cbr服务进行操作。 **约束限制**: 不涉及 **取值范围**: 字符长度0-64 **默认取值**: 不涉及
	CloseProtectionType string `json:"close_protection_type"`
}

func (o CloseProtectionInfoRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloseProtectionInfoRequestInfo struct{}"
	}

	return strings.Join([]string{"CloseProtectionInfoRequestInfo", string(data)}, " ")
}
