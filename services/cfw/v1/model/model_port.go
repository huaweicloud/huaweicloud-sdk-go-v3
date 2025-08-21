package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Port struct {

	// **参数解释**： 端口类型 **取值范围**： -1 Any，0 包含，1 排除
	PortType *int32 `json:"port_type,omitempty"`

	// **参数解释**： 端口 **取值范围**： 不涉及
	Ports *string `json:"ports,omitempty"`
}

func (o Port) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Port struct{}"
	}

	return strings.Join([]string{"Port", string(data)}, " ")
}
