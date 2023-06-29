package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetWebHookConfigRequest Request Object
type SetWebHookConfigRequest struct {
	Body *WebHookConfigRequest `json:"body,omitempty"`
}

func (o SetWebHookConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetWebHookConfigRequest struct{}"
	}

	return strings.Join([]string{"SetWebHookConfigRequest", string(data)}, " ")
}
