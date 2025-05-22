package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterConfigurationsParameterRequest Request Object
type ListClusterConfigurationsParameterRequest struct {

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// 参数组ID。
	ConfigurationId string `json:"configuration_id"`
}

func (o ListClusterConfigurationsParameterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterConfigurationsParameterRequest struct{}"
	}

	return strings.Join([]string{"ListClusterConfigurationsParameterRequest", string(data)}, " ")
}
