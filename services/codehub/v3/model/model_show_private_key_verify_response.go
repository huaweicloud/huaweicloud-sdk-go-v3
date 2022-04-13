package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowPrivateKeyVerifyResponse struct {
	Error *Error `json:"error,omitempty"`
	// 响应结果

	Result *string `json:"result,omitempty"`
	// 响应状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPrivateKeyVerifyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateKeyVerifyResponse struct{}"
	}

	return strings.Join([]string{"ShowPrivateKeyVerifyResponse", string(data)}, " ")
}
