package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableScheduledTaskResponse Response Object
type EnableScheduledTaskResponse struct {

	// 响应数据
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o EnableScheduledTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableScheduledTaskResponse struct{}"
	}

	return strings.Join([]string{"EnableScheduledTaskResponse", string(data)}, " ")
}
