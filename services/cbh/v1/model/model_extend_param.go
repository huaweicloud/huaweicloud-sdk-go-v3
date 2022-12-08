package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 安全组
type ExtendParam struct {

	// 计费模式 0:按需 1:包年包月
	ChargingMode string `json:"charging_mode"`

	// 磁盘资源类型
	RegionId string `json:"region_id"`
}

func (o ExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendParam struct{}"
	}

	return strings.Join([]string{"ExtendParam", string(data)}, " ")
}
