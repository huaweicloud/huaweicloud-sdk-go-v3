package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAiGuardLtsConfigRequest Request Object
type ShowAiGuardLtsConfigRequest struct {
}

func (o ShowAiGuardLtsConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAiGuardLtsConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowAiGuardLtsConfigRequest", string(data)}, " ")
}
