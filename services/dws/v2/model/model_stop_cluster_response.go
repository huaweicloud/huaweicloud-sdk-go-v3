package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopClusterResponse Response Object
type StopClusterResponse struct {

	// **参数解释**： 停止集群任务ID。 **取值范围**： 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopClusterResponse struct{}"
	}

	return strings.Join([]string{"StopClusterResponse", string(data)}, " ")
}
