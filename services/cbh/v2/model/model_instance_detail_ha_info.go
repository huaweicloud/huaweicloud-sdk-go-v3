package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceDetailHaInfo 主备信息。
type InstanceDetailHaInfo struct {

	// 主备ID。
	HaId string `json:"ha_id"`

	// 实例类型。 - master：主 - slave：备
	InstanceType string `json:"instance_type"`
}

func (o InstanceDetailHaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceDetailHaInfo struct{}"
	}

	return strings.Join([]string{"InstanceDetailHaInfo", string(data)}, " ")
}
