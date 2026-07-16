package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNetworkRequest Request Object
type DeleteNetworkRequest struct {

	// **参数解释**：网络资源名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	NetworkName string `json:"network_name"`
}

func (o DeleteNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNetworkRequest struct{}"
	}

	return strings.Join([]string{"DeleteNetworkRequest", string(data)}, " ")
}
