package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWdrSnapshotResponse Response Object
type CreateWdrSnapshotResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CreateWdrSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWdrSnapshotResponse struct{}"
	}

	return strings.Join([]string{"CreateWdrSnapshotResponse", string(data)}, " ")
}
