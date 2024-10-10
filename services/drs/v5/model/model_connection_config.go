package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConnectionConfig struct {

	// 驱动程序名称。
	DriverName *string `json:"driver_name,omitempty"`
}

func (o ConnectionConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionConfig struct{}"
	}

	return strings.Join([]string{"ConnectionConfig", string(data)}, " ")
}
