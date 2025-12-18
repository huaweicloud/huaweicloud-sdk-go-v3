package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InPlaceRollingUpdate 原地升级配置
type InPlaceRollingUpdate struct {

	// 节点升级步长，取值范围为[1, 40]，建议取值20
	UserDefinedStep *int32 `json:"userDefinedStep,omitempty"`

	// **参数解释：** 节点升级批次作用域 **约束限制：** 不涉及 **取值范围：** \"Cluster\"：节点升级批次配置应用到整个集群，整个升级过程不重置升级批次 \"NodePool\"：节点升级批次配置应用到节点池，升级每个节点池都会重置升级批次 **默认取值：** \"Cluster\"
	Scope *string `json:"scope,omitempty"`
}

func (o InPlaceRollingUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InPlaceRollingUpdate struct{}"
	}

	return strings.Join([]string{"InPlaceRollingUpdate", string(data)}, " ")
}
