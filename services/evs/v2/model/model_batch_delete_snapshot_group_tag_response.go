package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSnapshotGroupTagResponse Response Object
type BatchDeleteSnapshotGroupTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteSnapshotGroupTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSnapshotGroupTagResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteSnapshotGroupTagResponse", string(data)}, " ")
}
