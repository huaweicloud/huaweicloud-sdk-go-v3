package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunInstancesResponse Response Object
type RunInstancesResponse struct {

	// 实例返回信息
	Instances      *[]ServersResponsePhysicalServers `json:"instances,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o RunInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunInstancesResponse struct{}"
	}

	return strings.Join([]string{"RunInstancesResponse", string(data)}, " ")
}
