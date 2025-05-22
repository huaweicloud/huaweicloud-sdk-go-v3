package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterWorkloadRequest Request Object
type CreateClusterWorkloadRequest struct {

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	Body *WorkloadStatusReq `json:"body,omitempty"`
}

func (o CreateClusterWorkloadRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterWorkloadRequest struct{}"
	}

	return strings.Join([]string{"CreateClusterWorkloadRequest", string(data)}, " ")
}
