package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddonInstanceRollbackRequest **参数解释**： 回滚插件实例请求结构体。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type AddonInstanceRollbackRequest struct {

	// **参数解释**： 集群ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ClusterID string `json:"clusterID"`
}

func (o AddonInstanceRollbackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddonInstanceRollbackRequest struct{}"
	}

	return strings.Join([]string{"AddonInstanceRollbackRequest", string(data)}, " ")
}
