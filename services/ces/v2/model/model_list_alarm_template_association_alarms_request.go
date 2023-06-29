package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmTemplateAssociationAlarmsRequest Request Object
type ListAlarmTemplateAssociationAlarmsRequest struct {

	// 告警模板的ID，以at开头，后跟字母、数字，长度最长为64
	TemplateId string `json:"template_id"`

	// 分页查询时查询的起始位置，表示从第几条数据开始，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 查询结果条数的限制值，取值范围为[1,100]，默认值为100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAlarmTemplateAssociationAlarmsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmTemplateAssociationAlarmsRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmTemplateAssociationAlarmsRequest", string(data)}, " ")
}
