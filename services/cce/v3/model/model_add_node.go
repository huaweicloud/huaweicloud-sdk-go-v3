package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddNode 纳管节点参数。集群内已有节点支持通过重置进行重新安装并接入集群。
type AddNode struct {

	// **参数解释**： 服务器ID，从ECS/BMS控制台获取。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ServerID string `json:"serverID"`

	Spec *ReinstallNodeSpec `json:"spec"`
}

func (o AddNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddNode struct{}"
	}

	return strings.Join([]string{"AddNode", string(data)}, " ")
}
