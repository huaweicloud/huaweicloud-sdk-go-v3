package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterIdRes **参数解释**: 集群ID **取值范围**: 字符长度1-64位
type ClusterIdRes struct {
}

func (o ClusterIdRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterIdRes struct{}"
	}

	return strings.Join([]string{"ClusterIdRes", string(data)}, " ")
}
