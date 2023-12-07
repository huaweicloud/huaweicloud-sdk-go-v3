package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EgTargetDetail struct {

	// 目标项目id
	TargetProjectId string `json:"target_project_id"`

	// 目标通道id
	TargetChannelId string `json:"target_channel_id"`

	// 目标region
	TargetRegion string `json:"target_region"`

	// 跨region开关
	CrossRegion *bool `json:"cross_region,omitempty"`

	// 跨账号开关
	CrossAccount *bool `json:"cross_account,omitempty"`

	// 委托名称
	AgencyName string `json:"agency_name"`
}

func (o EgTargetDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EgTargetDetail struct{}"
	}

	return strings.Join([]string{"EgTargetDetail", string(data)}, " ")
}
