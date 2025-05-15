package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateP2PSiteNetworkRequestBody 创建P2P类型的分支网络的请求体。
type CreateP2PSiteNetworkRequestBody struct {
	P2pSiteNetwork *CreateP2PSiteNetwork `json:"p2p_site_network"`
}

func (o CreateP2PSiteNetworkRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateP2PSiteNetworkRequestBody struct{}"
	}

	return strings.Join([]string{"CreateP2PSiteNetworkRequestBody", string(data)}, " ")
}
