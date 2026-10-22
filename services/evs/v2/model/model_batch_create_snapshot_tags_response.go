package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateSnapshotTagsResponse Response Object
type BatchCreateSnapshotTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateSnapshotTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSnapshotTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateSnapshotTagsResponse", string(data)}, " ")
}
