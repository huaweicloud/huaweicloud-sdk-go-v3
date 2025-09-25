package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSlowLogDetailResponse Response Object
type ShowSlowLogDetailResponse struct {

	// **参数解释**：            慢日志列表。
	SlowLogList    *[]ShowStarRocksSlowLogDetail `json:"slow_log_list,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ShowSlowLogDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowLogDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowSlowLogDetailResponse", string(data)}, " ")
}
