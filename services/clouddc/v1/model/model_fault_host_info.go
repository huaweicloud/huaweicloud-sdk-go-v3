package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FaultHostInfo struct {

	// 告警时间
	FaultDate *sdktime.SdkTime `json:"fault_date,omitempty"`

	// 新增故障机器数
	NewFaultHost *int32 `json:"new_fault_host,omitempty"`

	// 故障机器SN
	FaultHost *[]string `json:"fault_host,omitempty"`
}

func (o FaultHostInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FaultHostInfo struct{}"
	}

	return strings.Join([]string{"FaultHostInfo", string(data)}, " ")
}
