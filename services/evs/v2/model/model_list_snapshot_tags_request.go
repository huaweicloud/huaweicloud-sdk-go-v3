package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotTagsRequest Request Object
type ListSnapshotTagsRequest struct {
}

func (o ListSnapshotTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotTagsRequest struct{}"
	}

	return strings.Join([]string{"ListSnapshotTagsRequest", string(data)}, " ")
}
