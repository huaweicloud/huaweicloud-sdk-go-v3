package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteBlockchainRequest struct {

	// bcs 服务id
	BlockchainId string `json:"blockchain_id"`

	// 是否删除存储，IEF模式下不用填写
	IsDeleteStorage *bool `json:"is_delete_storage,omitempty"`

	// 是否删除obs，IEF模式下不用填写
	IsDeleteObs *bool `json:"is_delete_obs,omitempty"`

	// 是否删除底层CCE资源，IEF模式下不用填写
	IsDeleteResource *bool `json:"is_delete_resource,omitempty"`

	// 是否删除底层依赖的IEF边缘节点资源，CCE模式下不用填写，IEF模式下必填
	IsDeleteIef *bool `json:"is_delete_ief,omitempty"`

	// CCE模式下选填：是否删除LightPeer轻节点插件使用的IEF边缘节点资源，若需要删除，则填写对应的IEF编译节点名称，否则无需填写。例如：is_delete_lightpeer=ief_node_name_1,ief_node_name_2
	IsDeleteLightpeer *string `json:"is_delete_lightpeer,omitempty"`

	// IEF模式下选填：若需要删除IEF边缘集群所使用的IEF边缘节点资源，则填写IEF节点的id，否则无需填写。例如：ief_nodes_id=1356f6f0-c448-4ec2-926f-c13b026369d9，0cdacbf9-2002-41e9-8123-2fa13e8d1449
	IefNodesId *string `json:"ief_nodes_id,omitempty"`
}

func (o DeleteBlockchainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBlockchainRequest struct{}"
	}

	return strings.Join([]string{"DeleteBlockchainRequest", string(data)}, " ")
}
