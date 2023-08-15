package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteAlarmTemplatesRequestBody struct {

	// 需要批量删除的告警模板的ID列表。未关联告警规则的模板可以批量删除多个；已关联告警规则的告警模板模板单次只允许删除一个，若同时删除多个已关联告警规则的告警模板，将返回异常
	TemplateIds []string `json:"template_ids"`

	// 如果告警模板关联了告警规则，是否级联删除告警规则，true代表级联删除，false代表只删除告警模板
	DeleteAssociateAlarm bool `json:"delete_associate_alarm"`
}

func (o BatchDeleteAlarmTemplatesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAlarmTemplatesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteAlarmTemplatesRequestBody", string(data)}, " ")
}
