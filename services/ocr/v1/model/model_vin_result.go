package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VinResult
type VinResult struct {

	// 识别检测到的车架号。
	Vin string `json:"vin"`
}

func (o VinResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VinResult struct{}"
	}

	return strings.Join([]string{"VinResult", string(data)}, " ")
}
