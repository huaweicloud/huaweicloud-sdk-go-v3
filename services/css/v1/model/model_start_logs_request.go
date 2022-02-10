package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StartLogsRequest struct {
	// 指定查询集群ID。

	ClusterId string `json:"cluster_id"`

	Body *StartLogsReq `json:"body,omitempty"`
}

func (o StartLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartLogsRequest struct{}"
	}

	return strings.Join([]string{"StartLogsRequest", string(data)}, " ")
}
