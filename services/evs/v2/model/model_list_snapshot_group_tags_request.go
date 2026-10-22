package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotGroupTagsRequest Request Object
type ListSnapshotGroupTagsRequest struct {
}

func (o ListSnapshotGroupTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotGroupTagsRequest struct{}"
	}

	return strings.Join([]string{"ListSnapshotGroupTagsRequest", string(data)}, " ")
}
