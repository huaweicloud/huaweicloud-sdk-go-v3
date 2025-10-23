package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VehicleLicenseAlarmResult
type VehicleLicenseAlarmResult struct {

	// -| 证件图像框内遮挡告警结果。 true：表示证件图片内部有被遮挡。 false：表示证件图片内部未被遮挡。
	DetectBlockingWithinBorderResult *bool `json:"detect_blocking_within_border_result,omitempty"`

	// -| 证件图像边框完整性告警结果。 true：表示边框不完整。 false：表示边框完整。
	DetectBorderIntegrityResult *bool `json:"detect_border_integrity_result,omitempty"`
}

func (o VehicleLicenseAlarmResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VehicleLicenseAlarmResult struct{}"
	}

	return strings.Join([]string{"VehicleLicenseAlarmResult", string(data)}, " ")
}
