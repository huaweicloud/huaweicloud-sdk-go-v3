package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SetWebHookConfigRequest struct {
	Body *WebHookConfigRequest `json:"body,omitempty" xml:"body"`
}

func (o SetWebHookConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetWebHookConfigRequest struct{}"
	}

	return strings.Join([]string{"SetWebHookConfigRequest", string(data)}, " ")
}
