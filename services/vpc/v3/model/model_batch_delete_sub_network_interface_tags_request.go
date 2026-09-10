package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSubNetworkInterfaceTagsRequest Request Object
type BatchDeleteSubNetworkInterfaceTagsRequest struct {

	// **参数解释**： 辅助弹性网卡唯一标识。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SubNetworkInterfaceId string `json:"sub_network_interface_id"`

	Body *BatchDeleteSubNetworkInterfaceTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteSubNetworkInterfaceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSubNetworkInterfaceTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteSubNetworkInterfaceTagsRequest", string(data)}, " ")
}
