package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EstimateDesktopPoolAddVolumeRequest Request Object
type EstimateDesktopPoolAddVolumeRequest struct {
	Body *EstimateAddVolumeRequestBody `json:"body,omitempty"`
}

func (o EstimateDesktopPoolAddVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EstimateDesktopPoolAddVolumeRequest struct{}"
	}

	return strings.Join([]string{"EstimateDesktopPoolAddVolumeRequest", string(data)}, " ")
}
