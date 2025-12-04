package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCoordinatorsRespCoordinators **参数解释**： 协调器信息。
type ShowCoordinatorsRespCoordinators struct {

	// **参数解释**： 消费组ID。 **取值范围**： 不涉及。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**： 对应协调器的broker id。 **取值范围**： 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释**： 对应协调器的地址。 **取值范围**： 不涉及。
	Host *string `json:"host,omitempty"`

	// **参数解释**： 端口号。 **取值范围**： 不涉及。
	Port *int32 `json:"port,omitempty"`
}

func (o ShowCoordinatorsRespCoordinators) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCoordinatorsRespCoordinators struct{}"
	}

	return strings.Join([]string{"ShowCoordinatorsRespCoordinators", string(data)}, " ")
}
