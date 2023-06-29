package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVmsSendTaskResponse Response Object
type CreateVmsSendTaskResponse struct {

	// 状态码。
	Status *string `json:"status,omitempty"`

	// 响应信息。
	Message *string `json:"message,omitempty"`

	Data           *CreateVmsTaskResponseMode `json:"data,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o CreateVmsSendTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVmsSendTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateVmsSendTaskResponse", string(data)}, " ")
}
