package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterEndpointsResponse Response Object
type ListClusterEndpointsResponse struct {
	PublicEndpoints *PublicEndpointResponse `json:"public_endpoints,omitempty"`

	PrivateEndpoints *PrivateEndpointResponse `json:"private_endpoints,omitempty"`

	// **参数解释**： 公网IP详细信息，显示每个节点当前是否绑定公网IP及对应状态。 **取值范围**： 不涉及。
	PublicIpInfos  *[]PublicIpInfoResponse `json:"public_ip_infos,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListClusterEndpointsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterEndpointsResponse struct{}"
	}

	return strings.Join([]string{"ListClusterEndpointsResponse", string(data)}, " ")
}
