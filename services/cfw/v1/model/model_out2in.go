package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Out2in struct {

	// **参数解释**： TOP访问目的IP **取值范围**： 不涉及
	DstIp *[]ItemVo `json:"dst_ip,omitempty"`

	// **参数解释**： TOP开放端口 **取值范围**： 不涉及
	DstPort *[]ItemVo `json:"dst_port,omitempty"`

	// **参数解释**： TOP访问源IP **取值范围**： 不涉及
	SrcIp *[]ItemVo `json:"src_ip,omitempty"`
}

func (o Out2in) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Out2in struct{}"
	}

	return strings.Join([]string{"Out2in", string(data)}, " ")
}
