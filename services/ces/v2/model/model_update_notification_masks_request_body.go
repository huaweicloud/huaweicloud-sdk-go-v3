package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNotificationMasksRequestBody 通知屏蔽请求体
type UpdateNotificationMasksRequestBody struct {

	// 屏蔽规则名称，只能为字母、数字、汉字、-、_，最大长度为64
	MaskName string `json:"mask_name"`

	// 关联编号，relation_type为ALARM_RULE时填屏蔽的告警规则编号；relation_type为RESOURCE_POLICY_NOTIFICATION、RESOURCE_POLICY_ALARM时填屏蔽的告警策略编号；
	RelationIds *[]string `json:"relation_ids,omitempty"`

	RelationType *RelationType `json:"relation_type,omitempty"`

	// 关联资源
	Resources []Resource `json:"resources"`

	MaskType *MaskType `json:"mask_type"`

	// 屏蔽起始日期，yyyy-MM-dd。
	StartDate *string `json:"start_date,omitempty"`

	// 屏蔽起始时间，HH:mm:ss。
	StartTime *string `json:"start_time,omitempty"`

	// 屏蔽截止日期，yyyy-MM-dd。
	EndDate *string `json:"end_date,omitempty"`

	// 屏蔽截止时间，HH:mm:ss。
	EndTime *string `json:"end_time,omitempty"`
}

func (o UpdateNotificationMasksRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotificationMasksRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateNotificationMasksRequestBody", string(data)}, " ")
}
