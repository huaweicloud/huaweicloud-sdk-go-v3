package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableControlParameters 策略参数。
type EnableControlParameters struct {

	// 策略参数名称。
	Key string `json:"key"`

	// 策略参数的值。
	Value *interface{} `json:"value"`
}

func (o EnableControlParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableControlParameters struct{}"
	}

	return strings.Join([]string{"EnableControlParameters", string(data)}, " ")
}
