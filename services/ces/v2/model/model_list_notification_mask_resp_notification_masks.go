package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListNotificationMaskRespNotificationMasks struct {

	// 屏蔽规则ID
	NotificationMaskId string `json:"notification_mask_id"`

	// 屏蔽规则名称，只能为字母、数字、汉字、-、_，最大长度为64
	MaskName *string `json:"mask_name,omitempty"`

	RelationType *RelationType `json:"relation_type"`

	// 关联编号
	RelationId *string `json:"relation_id,omitempty"`

	// 关联资源类型，relation_type为RESOURCE时存在该字段,只需要查询出资源的namespace+维度名即可
	Resources *[]ResourceCategory `json:"resources,omitempty"`

	MaskStatus *MaskStatus `json:"mask_status"`

	MaskType *MaskType `json:"mask_type"`

	// 屏蔽起始日期，yyyy-MM-dd。
	StartDate *string `json:"start_date,omitempty"`

	// 屏蔽起始时间，HH:mm:ss。
	StartTime *string `json:"start_time,omitempty"`

	// 屏蔽截止日期，yyyy-MM-dd。
	EndDate *string `json:"end_date,omitempty"`

	// 屏蔽截止时间，HH:mm:ss。
	EndTime *string `json:"end_time,omitempty"`

	// 告警策略列表。
	Policies *[]PoliciesInListResp `json:"policies,omitempty"`
}

func (o ListNotificationMaskRespNotificationMasks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationMaskRespNotificationMasks struct{}"
	}

	return strings.Join([]string{"ListNotificationMaskRespNotificationMasks", string(data)}, " ")
}
