package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchExceptionSnapshotStatusResponse Response Object
type SwitchExceptionSnapshotStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchExceptionSnapshotStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchExceptionSnapshotStatusResponse struct{}"
	}

	return strings.Join([]string{"SwitchExceptionSnapshotStatusResponse", string(data)}, " ")
}
