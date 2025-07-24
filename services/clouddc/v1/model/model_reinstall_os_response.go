package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReinstallOsResponse Response Object
type ReinstallOsResponse struct {

	// 实例返回信息
	Instances      *[]ServersResponsePhysicalServers `json:"instances,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ReinstallOsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallOsResponse struct{}"
	}

	return strings.Join([]string{"ReinstallOsResponse", string(data)}, " ")
}
