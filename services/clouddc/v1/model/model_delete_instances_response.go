package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstancesResponse Response Object
type DeleteInstancesResponse struct {

	// 实例返回信息
	Instances      *[]ServersResponsePhysicalServers `json:"instances,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o DeleteInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstancesResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstancesResponse", string(data)}, " ")
}
