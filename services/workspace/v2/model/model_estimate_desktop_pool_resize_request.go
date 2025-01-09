package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EstimateDesktopPoolResizeRequest Request Object
type EstimateDesktopPoolResizeRequest struct {
	Body *EstimateResizeRequestBody `json:"body,omitempty"`
}

func (o EstimateDesktopPoolResizeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EstimateDesktopPoolResizeRequest struct{}"
	}

	return strings.Join([]string{"EstimateDesktopPoolResizeRequest", string(data)}, " ")
}
