package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreClusterResponse Response Object
type RestoreClusterResponse struct {
	Cluster *Cluster `json:"cluster,omitempty"`

	// **参数解释**： 异步任务ID。 **取值范围**： 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreClusterResponse struct{}"
	}

	return strings.Join([]string{"RestoreClusterResponse", string(data)}, " ")
}
