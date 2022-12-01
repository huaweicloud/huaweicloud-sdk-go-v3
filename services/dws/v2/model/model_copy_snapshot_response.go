package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CopySnapshotResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CopySnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopySnapshotResponse struct{}"
	}

	return strings.Join([]string{"CopySnapshotResponse", string(data)}, " ")
}
