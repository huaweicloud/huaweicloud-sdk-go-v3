package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Fan 风扇的摘要信息，如型号、厂商等
type Fan struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 型号
	Model *string `json:"model,omitempty"`

	// 转速
	Reading *string `json:"reading,omitempty"`

	// 转速单位
	ReadingUnits *string `json:"reading_units,omitempty"`

	// 固件编码
	PartNumber *string `json:"part_number,omitempty"`

	// 速率比
	SpeedRatio *string `json:"speed_ratio,omitempty"`

	Status *Status `json:"status,omitempty"`
}

func (o Fan) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Fan struct{}"
	}

	return strings.Join([]string{"Fan", string(data)}, " ")
}
