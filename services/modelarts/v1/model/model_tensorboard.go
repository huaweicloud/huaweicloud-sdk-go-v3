package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Tensorboard Tensorboard连接信息。
type Tensorboard struct {

	// **参数解释**：训练作业的Tensorboard地址。 **取值范围**：不涉及。
	Url *string `json:"url,omitempty"`

	// **参数解释**：训练作业的Tensorboard token。 **取值范围**：不涉及。
	Token *string `json:"token,omitempty"`
}

func (o Tensorboard) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tensorboard struct{}"
	}

	return strings.Join([]string{"Tensorboard", string(data)}, " ")
}
