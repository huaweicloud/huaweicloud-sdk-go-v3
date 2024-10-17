package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CentralNetworkConnection 中心网络连接。
type CentralNetworkConnection struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例名字。
	Name string `json:"name"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 实例所属账号ID。
	DomainId string `json:"domain_id"`

	// 实例所属企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 中心网络ID。
	CentralNetworkId string `json:"central_network_id"`

	// 中心网络平面ID。
	CentralNetworkPlaneId string `json:"central_network_plane_id"`

	// 全域互联带宽ID。
	GlobalConnectionBandwidthId *string `json:"global_connection_bandwidth_id,omitempty"`

	BandwidthType *BandwidthTypeEnum `json:"bandwidth_type"`

	// 带宽值，单位Mbps。
	BandwidthSize *int64 `json:"bandwidth_size,omitempty"`

	State *CentralNetworkConnectionStateEnum `json:"state"`

	// 是否冻结
	IsFrozen bool `json:"is_frozen"`

	ConnectionType *ConnectionTypeEnum `json:"connection_type"`

	// 中心网络连接的两个端点定义，长度固定为2的数组。
	ConnectionPointPair []ConnectionPoint `json:"connection_point_pair"`

	// 实例创建时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 实例更新时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
}

func (o CentralNetworkConnection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CentralNetworkConnection struct{}"
	}

	return strings.Join([]string{"CentralNetworkConnection", string(data)}, " ")
}
