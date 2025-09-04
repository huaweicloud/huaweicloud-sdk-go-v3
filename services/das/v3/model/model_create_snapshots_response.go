package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSnapshotsResponse Response Object
type CreateSnapshotsResponse struct {

	// 是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreateSnapshotsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSnapshotsResponse struct{}"
	}

	return strings.Join([]string{"CreateSnapshotsResponse", string(data)}, " ")
}
