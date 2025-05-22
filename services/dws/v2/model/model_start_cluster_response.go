package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartClusterResponse Response Object
type StartClusterResponse struct {

	// **参数解释**： 启动集群的任务ID。 **取值范围**： 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartClusterResponse struct{}"
	}

	return strings.Join([]string{"StartClusterResponse", string(data)}, " ")
}
