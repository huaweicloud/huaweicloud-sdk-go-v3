package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeDbTemplateOperationResponse Response Object
type ChangeDbTemplateOperationResponse struct {

	// 返回消息
	Msg *string `json:"msg,omitempty"`

	// 返回状态，如'200','400'
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeDbTemplateOperationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeDbTemplateOperationResponse struct{}"
	}

	return strings.Join([]string{"ChangeDbTemplateOperationResponse", string(data)}, " ")
}
