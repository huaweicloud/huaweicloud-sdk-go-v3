package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HardwareSummary 硬件总览信息，包括服务器的制造商、型号、序列号等
type HardwareSummary struct {

	// serial number
	Sn string `json:"sn"`

	// 制造商
	Manufacturer string `json:"manufacturer"`

	// 型号
	Model string `json:"model"`

	// 主板厂商
	MainBoardManufacturer *string `json:"main_board_manufacturer,omitempty"`
}

func (o HardwareSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HardwareSummary struct{}"
	}

	return strings.Join([]string{"HardwareSummary", string(data)}, " ")
}
