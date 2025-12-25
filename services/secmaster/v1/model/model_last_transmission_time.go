package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LastTransmissionTime 最后一次传输的时间
type LastTransmissionTime struct {
}

func (o LastTransmissionTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LastTransmissionTime struct{}"
	}

	return strings.Join([]string{"LastTransmissionTime", string(data)}, " ")
}
