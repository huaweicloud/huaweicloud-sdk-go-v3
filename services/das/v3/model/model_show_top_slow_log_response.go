package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTopSlowLogResponse Response Object
type ShowTopSlowLogResponse struct {

	// TOP慢SQL列表
	TopSlowLogList *[]TopSlowLogInfo `json:"top_slow_log_list,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowTopSlowLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopSlowLogResponse struct{}"
	}

	return strings.Join([]string{"ShowTopSlowLogResponse", string(data)}, " ")
}
