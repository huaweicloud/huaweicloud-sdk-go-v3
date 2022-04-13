package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StartPublicWhitelistRequest struct {
	// 指定查询集群ID。

	ClusterId string `json:"cluster_id"`

	Body *StartPublicWhitelistReq `json:"body,omitempty"`
}

func (o StartPublicWhitelistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartPublicWhitelistRequest struct{}"
	}

	return strings.Join([]string{"StartPublicWhitelistRequest", string(data)}, " ")
}
