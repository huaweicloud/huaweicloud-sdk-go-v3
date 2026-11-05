package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopSupabaseApplicationResponse Response Object
type StopSupabaseApplicationResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopSupabaseApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopSupabaseApplicationResponse struct{}"
	}

	return strings.Join([]string{"StopSupabaseApplicationResponse", string(data)}, " ")
}
