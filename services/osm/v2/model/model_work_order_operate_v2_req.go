package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkOrderOperateV2Req struct {

	// 评价内容
	Judgement *string `json:"judgement,omitempty" xml:"judgement"`

	// 操作描述
	OperateDesc *string `json:"operate_desc,omitempty" xml:"operate_desc"`

	// 组id
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 工单满意度列表
	IncidentSatisfactionList *[]IncidentSatisfactionV2Do `json:"incident_satisfaction_list,omitempty" xml:"incident_satisfaction_list"`
}

func (o WorkOrderOperateV2Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkOrderOperateV2Req struct{}"
	}

	return strings.Join([]string{"WorkOrderOperateV2Req", string(data)}, " ")
}
