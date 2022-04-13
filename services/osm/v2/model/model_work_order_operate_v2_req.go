package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkOrderOperateV2Req struct {
	// 评价内容

	Judgement *string `json:"judgement,omitempty"`
	// 操作描述

	OperateDesc *string `json:"operate_desc,omitempty"`
	// 组id

	GroupId *string `json:"group_id,omitempty"`
	// 工单满意度列表

	IncidentSatisfactionList *[]IncidentSatisfactionV2Do `json:"incident_satisfaction_list,omitempty"`
}

func (o WorkOrderOperateV2Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkOrderOperateV2Req struct{}"
	}

	return strings.Join([]string{"WorkOrderOperateV2Req", string(data)}, " ")
}
