package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchSslResponse Response Object
type SwitchSslResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// SSL开关状态。
	SslOption      *string `json:"ssl_option,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchSslResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSslResponse struct{}"
	}

	return strings.Join([]string{"SwitchSslResponse", string(data)}, " ")
}
