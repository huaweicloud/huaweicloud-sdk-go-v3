package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAssociatedResourceOpenStatusResponse Response Object
type ChangeAssociatedResourceOpenStatusResponse struct {

	// 开通状态
	Status *string `json:"status,omitempty"`

	// 操作失败的错误信息
	Errors         *[]ErrorInfo `json:"errors,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ChangeAssociatedResourceOpenStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAssociatedResourceOpenStatusResponse struct{}"
	}

	return strings.Join([]string{"ChangeAssociatedResourceOpenStatusResponse", string(data)}, " ")
}
