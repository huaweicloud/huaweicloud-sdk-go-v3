package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EstimateDesktopPoolChangeImageRequest Request Object
type EstimateDesktopPoolChangeImageRequest struct {
	Body *EstimateChangeImageRequestBody `json:"body,omitempty"`
}

func (o EstimateDesktopPoolChangeImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EstimateDesktopPoolChangeImageRequest struct{}"
	}

	return strings.Join([]string{"EstimateDesktopPoolChangeImageRequest", string(data)}, " ")
}
