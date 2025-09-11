package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectWdrSnapshotResponse Response Object
type CollectWdrSnapshotResponse struct {

	// **参数解释**： 任务ID。 **取值范围**： 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CollectWdrSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectWdrSnapshotResponse struct{}"
	}

	return strings.Join([]string{"CollectWdrSnapshotResponse", string(data)}, " ")
}
