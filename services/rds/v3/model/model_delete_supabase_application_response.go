package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSupabaseApplicationResponse Response Object
type DeleteSupabaseApplicationResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteSupabaseApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSupabaseApplicationResponse struct{}"
	}

	return strings.Join([]string{"DeleteSupabaseApplicationResponse", string(data)}, " ")
}
