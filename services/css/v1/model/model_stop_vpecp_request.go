package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StopVpecpRequest struct {
	// 指定待关闭的集群ID。

	ClusterId string `json:"cluster_id"`
}

func (o StopVpecpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopVpecpRequest struct{}"
	}

	return strings.Join([]string{"StopVpecpRequest", string(data)}, " ")
}
