package model

import (
	"encoding/json"

	"strings"
)

type ChangingTheDelayThresholdRequestBody struct {
	// 延时阈值，单位是KB。
	DelayThresholdInKilobytes int32 `json:"delay_threshold_in_kilobytes"`
}

func (o ChangingTheDelayThresholdRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ChangingTheDelayThresholdRequestBody struct{}"
	}

	return strings.Join([]string{"ChangingTheDelayThresholdRequestBody", string(data)}, " ")
}
