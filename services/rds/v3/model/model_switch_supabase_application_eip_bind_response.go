package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchSupabaseApplicationEipBindResponse Response Object
type SwitchSupabaseApplicationEipBindResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchSupabaseApplicationEipBindResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSupabaseApplicationEipBindResponse struct{}"
	}

	return strings.Join([]string{"SwitchSupabaseApplicationEipBindResponse", string(data)}, " ")
}
