package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPortRequest Request Object
type ShowPortRequest struct {

	// **参数解释**： 端口的资源ID，可以是弹性网卡的资源ID。 **取值范围**： 不涉及。
	PortId string `json:"port_id"`
}

func (o ShowPortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPortRequest struct{}"
	}

	return strings.Join([]string{"ShowPortRequest", string(data)}, " ")
}
