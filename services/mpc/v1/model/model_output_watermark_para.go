package model

import (
	"encoding/json"

	"strings"
)

type OutputWatermarkPara struct {
	// 水印时长
	TimeDuration *int32 `json:"time_duration,omitempty"`
}

func (o OutputWatermarkPara) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "OutputWatermarkPara struct{}"
	}

	return strings.Join([]string{"OutputWatermarkPara", string(data)}, " ")
}
