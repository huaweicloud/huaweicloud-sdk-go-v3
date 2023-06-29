package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KafkaBrokerInfo Kafka Server地址
type KafkaBrokerInfo struct {

	// IP
	Ip string `json:"ip"`

	// Port
	Port int32 `json:"port"`
}

func (o KafkaBrokerInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KafkaBrokerInfo struct{}"
	}

	return strings.Join([]string{"KafkaBrokerInfo", string(data)}, " ")
}
