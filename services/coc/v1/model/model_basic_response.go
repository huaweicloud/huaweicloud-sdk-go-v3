package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BasicResponse 基础响应对象
type BasicResponse struct {

	// error_code
	ErrorCode *string `json:"error_code,omitempty"`

	// error_msg
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o BasicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasicResponse struct{}"
	}

	return strings.Join([]string{"BasicResponse", string(data)}, " ")
}
