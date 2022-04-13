package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceAbstractReq struct {
	// 实例描述

	Description *string `json:"description,omitempty"`
	// '维护时间窗开始时间。时间格式为 xx:00:00，xx取值为02,06,10,14,18,22。'  '在这个时间段内，运维人员可以对该实例的节点进行维护操作。维护期间，业务可以正常使用，可能会发生闪断。维护操作通常几个月一次。'

	MaintainBegin *string `json:"maintain_begin,omitempty"`
	// '维护时间窗结束时间。时间格式为 xx:00:00，与维护时间窗开始时间相差4个小时。'  '在这个时间段内，运维人员可以对该实例的节点进行维护操作。维护期间，业务可以正常使用，可能会发生闪断。维护操作通常几个月一次'。

	MaintainEnd *string `json:"maintain_end,omitempty"`
}

func (o InstanceAbstractReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceAbstractReq struct{}"
	}

	return strings.Join([]string{"InstanceAbstractReq", string(data)}, " ")
}
