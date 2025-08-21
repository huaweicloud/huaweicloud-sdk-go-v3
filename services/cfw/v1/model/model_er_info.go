package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ErInfo struct {

	// **参数解释**： 企业路由器ID **取值范围**： 不涉及
	ErId *string `json:"er_id,omitempty"`

	// **参数解释**： 企业路由器名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`
}

func (o ErInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErInfo struct{}"
	}

	return strings.Join([]string{"ErInfo", string(data)}, " ")
}
