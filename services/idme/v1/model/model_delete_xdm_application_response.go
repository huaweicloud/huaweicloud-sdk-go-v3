package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteXdmApplicationResponse Response Object
type DeleteXdmApplicationResponse struct {

	// 返回结果。
	Result *string `json:"result,omitempty"`

	// 错误信息。
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteXdmApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteXdmApplicationResponse struct{}"
	}

	return strings.Join([]string{"DeleteXdmApplicationResponse", string(data)}, " ")
}
