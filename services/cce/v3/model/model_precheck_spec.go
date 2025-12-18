package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrecheckSpec struct {

	// **参数解释：** 集群ID **约束限制：** 不涉及 **取值范围：** 不涉及
	ClusterID *string `json:"clusterID,omitempty"`

	// **参数解释：** 集群版本，请填写当前集群的补丁版本，可登录控制台在总览页面进行查看 **约束限制：** 不涉及 **取值范围：** 不涉及
	ClusterVersion *string `json:"clusterVersion,omitempty"`

	// **参数解释：** 升级目标版本，如果填写大版本，则自动选择最新补丁版本 **约束限制：** 不涉及 **取值范围：** 高于集群当前版本的可用集群版本
	TargetVersion *string `json:"targetVersion,omitempty"`

	// **参数解释：** 跳过检查的项目列表 **约束限制：** 不涉及 **取值范围：** 不涉及
	SkippedCheckItemList *[]SkippedCheckItemList `json:"skippedCheckItemList,omitempty"`
}

func (o PrecheckSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrecheckSpec struct{}"
	}

	return strings.Join([]string{"PrecheckSpec", string(data)}, " ")
}
