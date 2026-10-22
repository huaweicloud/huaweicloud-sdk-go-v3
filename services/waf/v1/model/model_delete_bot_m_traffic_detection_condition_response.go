package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBotMTrafficDetectionConditionResponse Response Object
type DeleteBotMTrafficDetectionConditionResponse struct {

	// **参数解释：** 结果 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Result         *bool `json:"result,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o DeleteBotMTrafficDetectionConditionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBotMTrafficDetectionConditionResponse struct{}"
	}

	return strings.Join([]string{"DeleteBotMTrafficDetectionConditionResponse", string(data)}, " ")
}
