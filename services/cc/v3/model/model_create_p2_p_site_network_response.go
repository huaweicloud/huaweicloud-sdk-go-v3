package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateP2PSiteNetworkResponse Response Object
type CreateP2PSiteNetworkResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	SiteNetwork    *SiteNetworkEntry `json:"site_network"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateP2PSiteNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateP2PSiteNetworkResponse struct{}"
	}

	return strings.Join([]string{"CreateP2PSiteNetworkResponse", string(data)}, " ")
}
