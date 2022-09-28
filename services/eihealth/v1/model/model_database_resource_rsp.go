package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DatabaseResourceRsp struct {

	// 实例ID
	Id string `json:"id"`

	// 资源ID
	ResourceId string `json:"resource_id"`

	Spec *DatabaseSpecDto `json:"spec"`

	Disk *DatabaseDiskDto `json:"disk"`

	// 计费模式
	ChargeMode string `json:"charge_mode"`

	// 购买周期
	PeriodNum int32 `json:"period_num"`

	// 购买时间，UTC时间
	CreateTime string `json:"create_time"`

	Status *DatabaseStatusEnum `json:"status"`
}

func (o DatabaseResourceRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseResourceRsp struct{}"
	}

	return strings.Join([]string{"DatabaseResourceRsp", string(data)}, " ")
}
