package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterLabelInfo **参数解释**: 集群资源标签 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及
type ClusterLabelInfo struct {

	// **参数解释**: 集群资源标签的键名 **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	Key *string `json:"key,omitempty"`

	// **参数解释**: 集群资源标签的值 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	Value *string `json:"value,omitempty"`
}

func (o ClusterLabelInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterLabelInfo struct{}"
	}

	return strings.Join([]string{"ClusterLabelInfo", string(data)}, " ")
}
