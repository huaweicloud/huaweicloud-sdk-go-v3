package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBackgroundTaskResponse Response Object
type DeleteBackgroundTaskResponse struct {

	// 返回消息
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteBackgroundTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackgroundTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteBackgroundTaskResponse", string(data)}, " ")
}
