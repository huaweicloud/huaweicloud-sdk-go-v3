package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ParaErrorResponse struct {

	// 错误码。
	ErrorCode string `json:"error_code"`

	// 错误消息。
	ErrorMsg string `json:"error_msg"`
}

func (o ParaErrorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParaErrorResponse struct{}"
	}

	return strings.Join([]string{"ParaErrorResponse", string(data)}, " ")
}
