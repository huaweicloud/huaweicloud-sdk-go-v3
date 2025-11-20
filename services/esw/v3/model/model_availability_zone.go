package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AvailabilityZone 可用区id
type AvailabilityZone struct {
}

func (o AvailabilityZone) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailabilityZone struct{}"
	}

	return strings.Join([]string{"AvailabilityZone", string(data)}, " ")
}
