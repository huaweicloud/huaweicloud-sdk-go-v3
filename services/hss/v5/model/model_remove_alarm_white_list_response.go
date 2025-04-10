package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveAlarmWhiteListResponse Response Object
type RemoveAlarmWhiteListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveAlarmWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveAlarmWhiteListResponse struct{}"
	}

	return strings.Join([]string{"RemoveAlarmWhiteListResponse", string(data)}, " ")
}
