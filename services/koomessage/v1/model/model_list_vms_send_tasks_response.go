package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVmsSendTasksResponse Response Object
type ListVmsSendTasksResponse struct {

	// 状态码。
	Status *string `json:"status,omitempty"`

	// 响应信息。
	Message *string `json:"message,omitempty"`

	Data           *ListVmsSendTasksResponseMode `json:"data,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListVmsSendTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVmsSendTasksResponse struct{}"
	}

	return strings.Join([]string{"ListVmsSendTasksResponse", string(data)}, " ")
}
