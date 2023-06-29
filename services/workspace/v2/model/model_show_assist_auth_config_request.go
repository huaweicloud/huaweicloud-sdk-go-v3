package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAssistAuthConfigRequest Request Object
type ShowAssistAuthConfigRequest struct {
}

func (o ShowAssistAuthConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssistAuthConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowAssistAuthConfigRequest", string(data)}, " ")
}
