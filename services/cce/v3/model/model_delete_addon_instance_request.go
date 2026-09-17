package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAddonInstanceRequest Request Object
type DeleteAddonInstanceRequest struct {

	// **参数解释**： 插件实例ID。 **约束限制**： 不涉及 **取值范围**： UUID格式，长度范围1~255位。 **默认取值**： 不涉及
	Id string `json:"id"`

	// **参数解释**： 集群ID（废弃中），获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。 **约束限制**： 不涉及 **取值范围**： UUID格式 **默认取值**： 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o DeleteAddonInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAddonInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteAddonInstanceRequest", string(data)}, " ")
}
