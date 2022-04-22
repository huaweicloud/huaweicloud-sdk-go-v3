package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StopPublicKibanaWhitelistRequest struct {

	// 指定待停止的集群ID。
	ClusterId string `json:"cluster_id"`
}

func (o StopPublicKibanaWhitelistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopPublicKibanaWhitelistRequest struct{}"
	}

	return strings.Join([]string{"StopPublicKibanaWhitelistRequest", string(data)}, " ")
}
