package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyInstanceSslConfigResponse Response Object
type ModifyInstanceSslConfigResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 协议模式。
	TlsMode        *string `json:"tls_mode,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyInstanceSslConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyInstanceSslConfigResponse struct{}"
	}

	return strings.Join([]string{"ModifyInstanceSslConfigResponse", string(data)}, " ")
}
