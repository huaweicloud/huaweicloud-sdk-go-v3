package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLabelsResponse Response Object
type UpdateLabelsResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateLabelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLabelsResponse struct{}"
	}

	return strings.Join([]string{"UpdateLabelsResponse", string(data)}, " ")
}
