package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CleanAlarmsResponse Response Object
type CleanAlarmsResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CleanAlarmsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CleanAlarmsResponse struct{}"
	}

	return strings.Join([]string{"CleanAlarmsResponse", string(data)}, " ")
}
