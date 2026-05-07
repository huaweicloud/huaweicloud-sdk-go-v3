package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterName **参数解释**: 集群名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
type ClusterName struct {
}

func (o ClusterName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterName struct{}"
	}

	return strings.Join([]string{"ClusterName", string(data)}, " ")
}
