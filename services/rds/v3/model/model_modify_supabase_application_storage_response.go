package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifySupabaseApplicationStorageResponse Response Object
type ModifySupabaseApplicationStorageResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifySupabaseApplicationStorageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifySupabaseApplicationStorageResponse struct{}"
	}

	return strings.Join([]string{"ModifySupabaseApplicationStorageResponse", string(data)}, " ")
}
