package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeServerPowerStateResponse Response Object
type ChangeServerPowerStateResponse struct {

	// 物理服务器返回信息
	PhysicalServers *[]ServersResponsePhysicalServers `json:"physical_servers,omitempty"`
	HttpStatusCode  int                               `json:"-"`
}

func (o ChangeServerPowerStateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeServerPowerStateResponse struct{}"
	}

	return strings.Join([]string{"ChangeServerPowerStateResponse", string(data)}, " ")
}
