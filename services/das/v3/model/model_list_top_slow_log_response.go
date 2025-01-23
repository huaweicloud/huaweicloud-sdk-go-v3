package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTopSlowLogResponse Response Object
type ListTopSlowLogResponse struct {

	// TOP慢SQL列表
	TopSlowLogList *[]TopSlowLogTopSlowLogList `json:"top_slow_log_list,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListTopSlowLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopSlowLogResponse struct{}"
	}

	return strings.Join([]string{"ListTopSlowLogResponse", string(data)}, " ")
}
