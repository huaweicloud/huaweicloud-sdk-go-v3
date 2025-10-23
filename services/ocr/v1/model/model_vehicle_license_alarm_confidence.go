package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VehicleLicenseAlarmConfidence
type VehicleLicenseAlarmConfidence struct {

	// 证件图像框内遮挡告警分数，分数越高，证件图像框内遮挡的可能性越高。
	BlockingWithinBorderScore *int32 `json:"blocking_within_border_score,omitempty"`

	// 证件图像边框完整性告警分数，分数越高，证件图像边框不完整的可能性越高。
	BorderIntegrityScore *int32 `json:"border_integrity_score,omitempty"`
}

func (o VehicleLicenseAlarmConfidence) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VehicleLicenseAlarmConfidence struct{}"
	}

	return strings.Join([]string{"VehicleLicenseAlarmConfidence", string(data)}, " ")
}
