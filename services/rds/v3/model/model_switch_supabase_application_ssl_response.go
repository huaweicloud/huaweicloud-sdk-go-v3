package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchSupabaseApplicationSslResponse Response Object
type SwitchSupabaseApplicationSslResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchSupabaseApplicationSslResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSupabaseApplicationSslResponse struct{}"
	}

	return strings.Join([]string{"SwitchSupabaseApplicationSslResponse", string(data)}, " ")
}
