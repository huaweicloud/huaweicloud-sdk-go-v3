package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSchedulesResponse Response Object
type ListSchedulesResponse struct {
	Data           *PageDataScheduleVo `json:"data,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListSchedulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSchedulesResponse struct{}"
	}

	return strings.Join([]string{"ListSchedulesResponse", string(data)}, " ")
}
