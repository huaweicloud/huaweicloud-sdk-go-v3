package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerGroupStateResponse Response Object
type ShowServerGroupStateResponse struct {

	// 对应状态的服务器数量，参考ServerStatus。
	ApsStatus      map[string]int32 `json:"aps_status,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowServerGroupStateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerGroupStateResponse struct{}"
	}

	return strings.Join([]string{"ShowServerGroupStateResponse", string(data)}, " ")
}
