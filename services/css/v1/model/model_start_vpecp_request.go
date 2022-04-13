package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StartVpecpRequest struct {
	// 指定待开启的集群ID。

	ClusterId string `json:"cluster_id"`

	Body *StartVpecpReq `json:"body,omitempty"`
}

func (o StartVpecpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartVpecpRequest struct{}"
	}

	return strings.Join([]string{"StartVpecpRequest", string(data)}, " ")
}
