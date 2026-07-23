package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyJobScheduleResponse Response Object
type ModifyJobScheduleResponse struct {

	// 响应结果
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyJobScheduleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyJobScheduleResponse struct{}"
	}

	return strings.Join([]string{"ModifyJobScheduleResponse", string(data)}, " ")
}
