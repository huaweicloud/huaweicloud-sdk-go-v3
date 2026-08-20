package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSprintSnapshotsResponse Response Object
type CreateSprintSnapshotsResponse struct {

	// 返回状态。
	Status *string `json:"status,omitempty"`

	// 请求失败时的错误信息。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSprintSnapshotsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSprintSnapshotsResponse struct{}"
	}

	return strings.Join([]string{"CreateSprintSnapshotsResponse", string(data)}, " ")
}
