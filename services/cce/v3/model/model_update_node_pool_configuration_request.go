package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNodePoolConfigurationRequest Request Object
type UpdateNodePoolConfigurationRequest struct {

	// 集群ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ClusterId string `json:"cluster_id"`

	// **参数解释：** 节点池ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。 **约束限制：** 不涉及 **取值范围：** - 节点池ID：修改指定节点池配置管理参数 - master：修改集群配置中心的配置管理参数  **默认取值：** 不涉及
	NodepoolId string `json:"nodepool_id"`

	Body *UpdateClusterConfigurationsBody `json:"body,omitempty"`
}

func (o UpdateNodePoolConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNodePoolConfigurationRequest struct{}"
	}

	return strings.Join([]string{"UpdateNodePoolConfigurationRequest", string(data)}, " ")
}
