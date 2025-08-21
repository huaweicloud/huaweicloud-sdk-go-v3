package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLogsCountRespData **参数解释**： 聚合时间点 **取值范围**： 不涉及
type ShowLogsCountRespData struct {

	// **参数解释**： 聚合时间点 **取值范围**： 不涉及
	Count *int64 `json:"count,omitempty"`
}

func (o ShowLogsCountRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogsCountRespData struct{}"
	}

	return strings.Join([]string{"ShowLogsCountRespData", string(data)}, " ")
}
