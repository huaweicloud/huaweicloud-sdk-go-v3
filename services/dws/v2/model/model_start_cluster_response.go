package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartClusterResponse Response Object
type StartClusterResponse struct {

	// 启动集群jobId
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
