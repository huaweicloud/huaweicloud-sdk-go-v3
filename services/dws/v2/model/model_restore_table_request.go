package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreTableRequest Request Object
type RestoreTableRequest struct {

	// 快照ID
	SnapshotId string `json:"snapshot_id"`

	Body *RestoreTableRequestBody `json:"body,omitempty"`
}

func (o RestoreTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreTableRequest struct{}"
	}

	return strings.Join([]string{"RestoreTableRequest", string(data)}, " ")
}
