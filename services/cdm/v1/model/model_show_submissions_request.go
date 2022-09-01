package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowSubmissionsRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id" xml:"cluster_id"`

	// 作业名称
	Jname string `json:"jname" xml:"jname"`
}

func (o ShowSubmissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubmissionsRequest struct{}"
	}

	return strings.Join([]string{"ShowSubmissionsRequest", string(data)}, " ")
}
