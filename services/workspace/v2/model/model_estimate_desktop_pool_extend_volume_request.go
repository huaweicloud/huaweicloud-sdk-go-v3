package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EstimateDesktopPoolExtendVolumeRequest Request Object
type EstimateDesktopPoolExtendVolumeRequest struct {
	Body *EstimateExtendVolumeRequestBody `json:"body,omitempty"`
}

func (o EstimateDesktopPoolExtendVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EstimateDesktopPoolExtendVolumeRequest struct{}"
	}

	return strings.Join([]string{"EstimateDesktopPoolExtendVolumeRequest", string(data)}, " ")
}
