package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateSnapshotGroupTagResponse Response Object
type BatchCreateSnapshotGroupTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateSnapshotGroupTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSnapshotGroupTagResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateSnapshotGroupTagResponse", string(data)}, " ")
}
