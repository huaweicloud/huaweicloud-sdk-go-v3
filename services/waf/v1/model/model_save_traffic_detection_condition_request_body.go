package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SaveTrafficDetectionConditionRequestBody struct {
	Condition *TrafficDetectionConditionDto `json:"condition,omitempty"`
}

func (o SaveTrafficDetectionConditionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveTrafficDetectionConditionRequestBody struct{}"
	}

	return strings.Join([]string{"SaveTrafficDetectionConditionRequestBody", string(data)}, " ")
}
