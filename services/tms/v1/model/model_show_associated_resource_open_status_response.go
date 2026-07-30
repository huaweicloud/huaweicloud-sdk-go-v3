package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAssociatedResourceOpenStatusResponse Response Object
type ShowAssociatedResourceOpenStatusResponse struct {

	// 开通状态
	Status *string `json:"status,omitempty"`

	// 操作失败的错误信息
	Errors         *[]ErrorInfo `json:"errors,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowAssociatedResourceOpenStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssociatedResourceOpenStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowAssociatedResourceOpenStatusResponse", string(data)}, " ")
}
