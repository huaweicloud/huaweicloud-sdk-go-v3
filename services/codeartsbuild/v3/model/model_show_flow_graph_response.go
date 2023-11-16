package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlowGraphResponse Response Object
type ShowFlowGraphResponse struct {

	// 状态
	Success *bool `json:"success,omitempty"`

	// 消息
	Message *string `json:"message,omitempty"`

	// 错误码
	ErrCode *string `json:"err_code,omitempty"`

	Result         *FlowGraphResult `json:"result,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowFlowGraphResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlowGraphResponse struct{}"
	}

	return strings.Join([]string{"ShowFlowGraphResponse", string(data)}, " ")
}
