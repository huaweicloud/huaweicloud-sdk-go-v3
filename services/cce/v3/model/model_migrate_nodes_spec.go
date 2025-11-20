package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MigrateNodesSpec struct {

	// **参数解释**： 操作系统类型，须精确到版本号。例：Huawei Cloud EulerOS 2.0。具体支持的操作系统请参见[节点操作系统说明](node-os.xml)。 **约束限制**： 当指定“alpha.cce/NodeImageID”参数时，“os”参数必须和用户自定义镜像的操作系统一致。 **取值范围**： 不涉及 **默认取值**： 不涉及
	Os string `json:"os"`

	ExtendParam *MigrateNodeExtendParam `json:"extendParam,omitempty"`

	Login *Login `json:"login,omitempty"`

	Runtime *Runtime `json:"runtime,omitempty"`

	ServerConfig *MigrateServerConfig `json:"serverConfig,omitempty"`

	// **参数解释**： 待操作节点列表，当前最多支持同时迁移200个节点。 **约束限制**： 不涉及
	Nodes []NodeItem `json:"nodes"`
}

func (o MigrateNodesSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateNodesSpec struct{}"
	}

	return strings.Join([]string{"MigrateNodesSpec", string(data)}, " ")
}
