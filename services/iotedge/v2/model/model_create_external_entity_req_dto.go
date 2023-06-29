package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateExternalEntityReqDto 创建外部实体请求结构体
type CreateExternalEntityReqDto struct {

	// 外部实体Id，节点下唯一
	ExternalId string `json:"external_id"`

	// 连接外部实体的协议类型
	Protocol string `json:"protocol"`

	// 连接类型
	ConnectionType string `json:"connection_type"`

	MqttConnectionInfo *MqttConnectionInfo `json:"mqtt_connection_info,omitempty"`

	// 资源空间ID。此参数为非必选参数，存在多资源空间的用户需要使用该接口时，建议携带该参数指定创建的外部实体归属到IoDA哪个资源空间下的边缘节点设备下，否则创建的外部实体将会归属到默认资源空间下对应的边缘节点下,对应于IoDA的app_id.
	SpaceId *string `json:"space_id,omitempty"`
}

func (o CreateExternalEntityReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExternalEntityReqDto struct{}"
	}

	return strings.Join([]string{"CreateExternalEntityReqDto", string(data)}, " ")
}
