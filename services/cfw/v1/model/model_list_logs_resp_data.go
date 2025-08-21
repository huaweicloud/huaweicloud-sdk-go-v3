package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogsRespData **参数解释**： 日志 **取值范围**： 不涉及
type ListLogsRespData struct {

	// **参数解释**： 条数 **取值范围**： 不涉及
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 日志 **取值范围**： 不涉及
	Records *[]LogVo `json:"records,omitempty"`

	// **参数解释**： 记录总数 **取值范围**： 不涉及
	Total *int64 `json:"total,omitempty"`
}

func (o ListLogsRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogsRespData struct{}"
	}

	return strings.Join([]string{"ListLogsRespData", string(data)}, " ")
}
