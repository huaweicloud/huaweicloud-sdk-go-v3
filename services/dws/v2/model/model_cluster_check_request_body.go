package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterCheckRequestBody **参数解释**： 创建集群校验请求体。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type ClusterCheckRequestBody struct {
	Cluster *ClusterCheckBody `json:"cluster"`
}

func (o ClusterCheckRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterCheckRequestBody struct{}"
	}

	return strings.Join([]string{"ClusterCheckRequestBody", string(data)}, " ")
}
