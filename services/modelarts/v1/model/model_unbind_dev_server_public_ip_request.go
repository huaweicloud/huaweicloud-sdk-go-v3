package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnbindDevServerPublicIpRequest Request Object
type UnbindDevServerPublicIpRequest struct {

	// **参数解释**：Lite Server实例ID。 **默认取值**：不涉及。
	Id string `json:"id"`

	// **参数解释**：EIP ID。 **默认取值**：不涉及。
	PublicipId string `json:"publicip_id"`
}

func (o UnbindDevServerPublicIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindDevServerPublicIpRequest struct{}"
	}

	return strings.Join([]string{"UnbindDevServerPublicIpRequest", string(data)}, " ")
}
