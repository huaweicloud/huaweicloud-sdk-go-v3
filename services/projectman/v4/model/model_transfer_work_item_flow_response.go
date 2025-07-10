package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransferWorkItemFlowResponse Response Object
type TransferWorkItemFlowResponse struct {

	// 状态码
	Status *string `json:"status,omitempty"`

	// 响应信息
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o TransferWorkItemFlowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferWorkItemFlowResponse struct{}"
	}

	return strings.Join([]string{"TransferWorkItemFlowResponse", string(data)}, " ")
}
