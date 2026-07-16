package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeInfo 节点检查状态
type NodeInfo struct {

	// 节点UID
	Uid *string `json:"uid,omitempty"`

	// 节点名称
	Name *string `json:"name,omitempty"`

	// **参数解释**： 节点状态 **取值范围**： - Build：创建中，表示节点正处于创建过程中。 - Installing：安装中，表示节点正处于纳管过程中。 - Upgrading：升级中，表示节点正处于升级过程中。 - Active：运行中，表示节点处于正常状态。 - Abnormal：不可用，表示节点处于异常状态。 - Deleting： 删除中，表示节点正处于删除过程中。 - Error：错误，表示节点处于故障状态。 **默认取值**： 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**： 节点类型 **取值范围**： - master：控制面节点 - node：数据面节点 **默认取值**： 不涉及
	NodeType *string `json:"nodeType,omitempty"`
}

func (o NodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeInfo struct{}"
	}

	return strings.Join([]string{"NodeInfo", string(data)}, " ")
}
