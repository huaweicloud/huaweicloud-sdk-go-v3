package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAiAppRequest Request Object
type DeleteAiAppRequest struct {

	// **参数解释：** AI应用ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AppId string `json:"app_id"`

	// **参数解释：** 是否保留关联的策略 **约束限制：** 不涉及 **取值范围：** true, false **默认取值：** 不涉及
	KeepPolicy bool `json:"keep_policy"`
}

func (o DeleteAiAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAiAppRequest struct{}"
	}

	return strings.Join([]string{"DeleteAiAppRequest", string(data)}, " ")
}
