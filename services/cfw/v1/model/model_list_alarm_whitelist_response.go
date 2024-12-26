package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmWhitelistResponse Response Object
type ListAlarmWhitelistResponse struct {
	Data           *PageInfo `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAlarmWhitelistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmWhitelistResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmWhitelistResponse", string(data)}, " ")
}
