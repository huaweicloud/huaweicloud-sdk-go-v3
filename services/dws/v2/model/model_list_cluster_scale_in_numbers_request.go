package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterScaleInNumbersRequest Request Object
type ListClusterScaleInNumbersRequest struct {

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`
}

func (o ListClusterScaleInNumbersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterScaleInNumbersRequest struct{}"
	}

	return strings.Join([]string{"ListClusterScaleInNumbersRequest", string(data)}, " ")
}
