package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateSubNetworkInterfaceTagsRequest Request Object
type BatchCreateSubNetworkInterfaceTagsRequest struct {

	// **参数解释**： 辅助弹性网卡唯一标识。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SubNetworkInterfaceId string `json:"sub_network_interface_id"`

	Body *BatchCreateSubNetworkInterfaceTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateSubNetworkInterfaceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSubNetworkInterfaceTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateSubNetworkInterfaceTagsRequest", string(data)}, " ")
}
