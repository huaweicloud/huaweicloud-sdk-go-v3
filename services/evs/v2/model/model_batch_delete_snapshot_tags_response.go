package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSnapshotTagsResponse Response Object
type BatchDeleteSnapshotTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteSnapshotTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSnapshotTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteSnapshotTagsResponse", string(data)}, " ")
}
