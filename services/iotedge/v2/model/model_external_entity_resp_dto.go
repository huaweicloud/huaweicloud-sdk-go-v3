package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建外部实体后返回的信息结构体
type ExternalEntityRespDto struct {

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 资源空间ID。此参数为非必选参数，存在多资源空间的用户需要使用该接口时，建议携带该参数指定创建的外部实体归属到IoDA哪个资源空间下的边缘节点设备下，否则创建的外部实体将会归属到默认资源空间下对应的边缘节点下,对应于IoDA的app_id.
	SpaceId *string `json:"space_id,omitempty"`

	// 外部实体ID
	ExternalId *string `json:"external_id,omitempty"`

	// 接入协议类型
	Protocol *string `json:"protocol,omitempty"`

	// 连接类型(client和server)
	ConnectionType *string `json:"connection_type,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 最后修改时间
	LastModifyTime *string `json:"last_modify_time,omitempty"`
}

func (o ExternalEntityRespDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalEntityRespDto struct{}"
	}

	return strings.Join([]string{"ExternalEntityRespDto", string(data)}, " ")
}
