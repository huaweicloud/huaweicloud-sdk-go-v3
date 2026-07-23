package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteJobScheduleResponse Response Object
type DeleteJobScheduleResponse struct {

	// 响应结果
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteJobScheduleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJobScheduleResponse struct{}"
	}

	return strings.Join([]string{"DeleteJobScheduleResponse", string(data)}, " ")
}
