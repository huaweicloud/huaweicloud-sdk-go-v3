package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 可用区。
type AvailabilityZone struct {
	// 可用区唯一编码。

	Code string `json:"code"`
	// 可用区状态。  取值：ACTIVE。

	State string `json:"state"`
}

func (o AvailabilityZone) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailabilityZone struct{}"
	}

	return strings.Join([]string{"AvailabilityZone", string(data)}, " ")
}
