package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetSupabaseApplicationApiKeyResponse Response Object
type ResetSupabaseApplicationApiKeyResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResetSupabaseApplicationApiKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetSupabaseApplicationApiKeyResponse struct{}"
	}

	return strings.Join([]string{"ResetSupabaseApplicationApiKeyResponse", string(data)}, " ")
}
