package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableLtsLogsRequest Request Object
type DisableLtsLogsRequest struct {

	// 集群的ID
	ClusterId string `json:"cluster_id"`
}

func (o DisableLtsLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableLtsLogsRequest struct{}"
	}

	return strings.Join([]string{"DisableLtsLogsRequest", string(data)}, " ")
}
