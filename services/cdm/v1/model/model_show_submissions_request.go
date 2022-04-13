package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowSubmissionsRequest struct {
	// 集群ID

	ClusterId string `json:"cluster_id"`
	// 作业名称

	Jname string `json:"jname"`
}

func (o ShowSubmissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubmissionsRequest struct{}"
	}

	return strings.Join([]string{"ShowSubmissionsRequest", string(data)}, " ")
}
