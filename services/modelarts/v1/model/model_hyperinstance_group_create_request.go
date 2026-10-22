package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HyperinstanceGroupCreateRequest struct {

	// **参数解释**：hyperinstance-group名称。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：超节点集群ID。 **默认取值**：不涉及。
	HyperinstanceClusterId string `json:"hyperinstance_cluster_id"`
}

func (o HyperinstanceGroupCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HyperinstanceGroupCreateRequest struct{}"
	}

	return strings.Join([]string{"HyperinstanceGroupCreateRequest", string(data)}, " ")
}
