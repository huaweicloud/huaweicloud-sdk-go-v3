package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotsRequest Request Object
type ListSnapshotsRequest struct {
}

func (o ListSnapshotsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotsRequest struct{}"
	}

	return strings.Join([]string{"ListSnapshotsRequest", string(data)}, " ")
}
