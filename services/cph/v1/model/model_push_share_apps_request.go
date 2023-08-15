package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PushShareAppsRequest Request Object
type PushShareAppsRequest struct {
	Body *PushShareAppsRequestBody `json:"body,omitempty"`
}

func (o PushShareAppsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PushShareAppsRequest struct{}"
	}

	return strings.Join([]string{"PushShareAppsRequest", string(data)}, " ")
}
