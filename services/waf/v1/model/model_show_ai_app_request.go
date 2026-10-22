package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAiAppRequest Request Object
type ShowAiAppRequest struct {

	// **参数解释：** AI应用ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AppId string `json:"app_id"`
}

func (o ShowAiAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAiAppRequest struct{}"
	}

	return strings.Join([]string{"ShowAiAppRequest", string(data)}, " ")
}
