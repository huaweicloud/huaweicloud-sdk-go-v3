package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifySupabaseApplicationPasswordResponse Response Object
type ModifySupabaseApplicationPasswordResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifySupabaseApplicationPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifySupabaseApplicationPasswordResponse struct{}"
	}

	return strings.Join([]string{"ModifySupabaseApplicationPasswordResponse", string(data)}, " ")
}
