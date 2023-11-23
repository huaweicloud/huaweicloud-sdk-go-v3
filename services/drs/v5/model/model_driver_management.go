package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DriverManagement struct {

	// 驱动名称。
	DriverName *string `json:"driver_name,omitempty"`
}

func (o DriverManagement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DriverManagement struct{}"
	}

	return strings.Join([]string{"DriverManagement", string(data)}, " ")
}
