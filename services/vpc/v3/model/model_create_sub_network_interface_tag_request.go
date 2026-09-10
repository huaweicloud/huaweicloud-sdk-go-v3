package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubNetworkInterfaceTagRequest Request Object
type CreateSubNetworkInterfaceTagRequest struct {

	// **参数解释**： 辅助弹性网卡唯一标识。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SubNetworkInterfaceId string `json:"sub_network_interface_id"`

	Body *CreateSubNetworkInterfaceTagRequestBody `json:"body,omitempty"`
}

func (o CreateSubNetworkInterfaceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubNetworkInterfaceTagRequest struct{}"
	}

	return strings.Join([]string{"CreateSubNetworkInterfaceTagRequest", string(data)}, " ")
}
