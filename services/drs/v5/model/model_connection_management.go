package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConnectionManagement struct {
	DriverManagement *DriverManagement `json:"driver_management,omitempty"`
}

func (o ConnectionManagement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionManagement struct{}"
	}

	return strings.Join([]string{"ConnectionManagement", string(data)}, " ")
}
