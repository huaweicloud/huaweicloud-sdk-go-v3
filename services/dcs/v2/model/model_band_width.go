package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BandWidth 带宽信息
type BandWidth struct {

	// 上行带宽，单位kbit/s
	Input *string `json:"input,omitempty"`

	// 下行带宽，单位kbit/s
	Output *string `json:"output,omitempty"`
}

func (o BandWidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandWidth struct{}"
	}

	return strings.Join([]string{"BandWidth", string(data)}, " ")
}
