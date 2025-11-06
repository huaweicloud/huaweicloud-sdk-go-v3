package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmsResponse Response Object
type ListAlarmsResponse struct {

	// **参数解释**： 查询的告警对象列表。
	MetricAlarms *[]MetricAlarmsResp `json:"metric_alarms,omitempty"`

	MetaData       *MetaDataResp `json:"meta_data,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListAlarmsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmsResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmsResponse", string(data)}, " ")
}
