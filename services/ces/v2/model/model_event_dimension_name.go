package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventDimensionName 事件维度名，多个维度按字母序逗号分开
type EventDimensionName struct {
}

func (o EventDimensionName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventDimensionName struct{}"
	}

	return strings.Join([]string{"EventDimensionName", string(data)}, " ")
}
