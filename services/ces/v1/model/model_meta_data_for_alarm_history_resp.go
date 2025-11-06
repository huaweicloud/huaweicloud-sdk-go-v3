package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetaDataForAlarmHistoryResp **参数解释** 查询告警历史返回的总条数。
type MetaDataForAlarmHistoryResp struct {

	// **参数解释** 查询告警历史返回的总条数。 **取值范围**： 不涉及。
	Total *int32 `json:"total,omitempty"`
}

func (o MetaDataForAlarmHistoryResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetaDataForAlarmHistoryResp struct{}"
	}

	return strings.Join([]string{"MetaDataForAlarmHistoryResp", string(data)}, " ")
}
