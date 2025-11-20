package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConnectionRequest Request Object
type ShowConnectionRequest struct {

	// - 参数解释：二层连接的唯一标识。 - 约束限制：带“-”的UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	ConnectionId string `json:"connection_id"`

	// - 参数解释：ESW实例的唯一标识。 - 约束限制：带“-”的UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	InstanceId string `json:"instance_id"`
}

func (o ShowConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConnectionRequest struct{}"
	}

	return strings.Join([]string{"ShowConnectionRequest", string(data)}, " ")
}
