package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScheduledTaskResponse Response Object
type CreateScheduledTaskResponse struct {

	// 响应数据
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateScheduledTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScheduledTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateScheduledTaskResponse", string(data)}, " ")
}
