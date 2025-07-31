package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Nameserver struct {

	// **参数解释：** 主机名。 **取值范围：** 不涉及。
	Hostname *string `json:"hostname,omitempty"`

	// **参数解释：** 优先级。 **取值范围：** 不涉及。
	Priority *int32 `json:"priority,omitempty"`
}

func (o Nameserver) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nameserver struct{}"
	}

	return strings.Join([]string{"Nameserver", string(data)}, " ")
}
