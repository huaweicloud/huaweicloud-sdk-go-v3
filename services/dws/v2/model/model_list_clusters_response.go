package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClustersResponse Response Object
type ListClustersResponse struct {

	// **参数解释**： 集群对象列表。 **取值范围**： 不涉及。
	Clusters *[]ClusterInfo `json:"clusters,omitempty"`

	// **参数解释**： 集群对象列表总数。 **取值范围**： 不涉及。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListClustersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClustersResponse struct{}"
	}

	return strings.Join([]string{"ListClustersResponse", string(data)}, " ")
}
