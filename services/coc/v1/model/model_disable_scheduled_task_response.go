package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableScheduledTaskResponse Response Object
type DisableScheduledTaskResponse struct {

	// 响应数据
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DisableScheduledTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableScheduledTaskResponse struct{}"
	}

	return strings.Join([]string{"DisableScheduledTaskResponse", string(data)}, " ")
}
