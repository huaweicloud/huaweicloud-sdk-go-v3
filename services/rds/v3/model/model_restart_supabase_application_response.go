package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartSupabaseApplicationResponse Response Object
type RestartSupabaseApplicationResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestartSupabaseApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartSupabaseApplicationResponse struct{}"
	}

	return strings.Join([]string{"RestartSupabaseApplicationResponse", string(data)}, " ")
}
