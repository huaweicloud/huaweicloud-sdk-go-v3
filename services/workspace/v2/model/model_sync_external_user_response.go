package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncExternalUserResponse Response Object
type SyncExternalUserResponse struct {

	// 同步用户（组）的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SyncExternalUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncExternalUserResponse struct{}"
	}

	return strings.Join([]string{"SyncExternalUserResponse", string(data)}, " ")
}
