package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PatchNetworkRequest Request Object
type PatchNetworkRequest struct {

	// **参数解释**：网络资源名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	NetworkName string `json:"network_name"`

	Body *NetworkUpdateRequest `json:"body,omitempty"`
}

func (o PatchNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PatchNetworkRequest struct{}"
	}

	return strings.Join([]string{"PatchNetworkRequest", string(data)}, " ")
}
