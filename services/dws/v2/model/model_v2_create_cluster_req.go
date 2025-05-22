package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// V2CreateClusterReq **参数解释**： V2接口创建集群请求体。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type V2CreateClusterReq struct {
	Cluster *V2CreateCluster `json:"cluster,omitempty"`
}

func (o V2CreateClusterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2CreateClusterReq struct{}"
	}

	return strings.Join([]string{"V2CreateClusterReq", string(data)}, " ")
}
