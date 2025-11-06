package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterVolumeUsageRequest Request Object
type ShowClusterVolumeUsageRequest struct {

	// **参数解释**： 指定更新的集群ID，获取方法请参见获取集群ID。 **约束限制**： 不涉及 **取值范围**： 获取方法请参见获取集群ID。 **默认取值**： 不涉及
	ClusterId string `json:"cluster_id"`
}

func (o ShowClusterVolumeUsageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterVolumeUsageRequest struct{}"
	}

	return strings.Join([]string{"ShowClusterVolumeUsageRequest", string(data)}, " ")
}
