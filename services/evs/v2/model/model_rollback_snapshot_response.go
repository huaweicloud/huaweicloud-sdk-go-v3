package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RollbackSnapshotResponse struct {
	Rollback       *RollbackInfo `json:"rollback,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o RollbackSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollbackSnapshotResponse struct{}"
	}

	return strings.Join([]string{"RollbackSnapshotResponse", string(data)}, " ")
}
