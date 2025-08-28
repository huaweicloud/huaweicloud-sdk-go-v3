package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLogtankRequest Request Object
type ShowLogtankRequest struct {

	// **参数解释**：云日志ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	LogtankId string `json:"logtank_id"`
}

func (o ShowLogtankRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogtankRequest struct{}"
	}

	return strings.Join([]string{"ShowLogtankRequest", string(data)}, " ")
}
