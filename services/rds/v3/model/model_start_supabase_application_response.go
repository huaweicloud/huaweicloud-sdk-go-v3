package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartSupabaseApplicationResponse Response Object
type StartSupabaseApplicationResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartSupabaseApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartSupabaseApplicationResponse struct{}"
	}

	return strings.Join([]string{"StartSupabaseApplicationResponse", string(data)}, " ")
}
