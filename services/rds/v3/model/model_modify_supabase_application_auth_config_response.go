package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifySupabaseApplicationAuthConfigResponse Response Object
type ModifySupabaseApplicationAuthConfigResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifySupabaseApplicationAuthConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifySupabaseApplicationAuthConfigResponse struct{}"
	}

	return strings.Join([]string{"ModifySupabaseApplicationAuthConfigResponse", string(data)}, " ")
}
