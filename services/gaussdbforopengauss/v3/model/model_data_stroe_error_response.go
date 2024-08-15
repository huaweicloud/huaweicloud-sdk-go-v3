package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataStroeErrorResponse struct {

	// 错误码。
	ErrorCode string `json:"error_code"`

	// 错误消息。
	ErrorMsg string `json:"error_msg"`
}

func (o DataStroeErrorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataStroeErrorResponse struct{}"
	}

	return strings.Join([]string{"DataStroeErrorResponse", string(data)}, " ")
}
