package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowScheduleResponse Response Object
type ShowScheduleResponse struct {

	// 是否可调度
	Schedulable    *bool `json:"schedulable,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowScheduleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScheduleResponse struct{}"
	}

	return strings.Join([]string{"ShowScheduleResponse", string(data)}, " ")
}
