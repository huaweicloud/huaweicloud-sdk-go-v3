package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterClusterSpec 集群spec信息。
type RegisterClusterSpec struct {

	// 容器舰队ID信息。
	ClusterGroupID *string `json:"clusterGroupID,omitempty"`

	// 集群类别，填写需要与provider和type对应，具体请参见[集群类别与类型说明](ucs_api_0024.xml)。
	Category string `json:"category"`

	// 集群类型，填写需要与category和provider对应，具体请参见[集群类别与类型说明](ucs_api_0024.xml)。
	Type string `json:"type"`

	// 供应商，填写需要与category和type对应，具体请参见[集群类别与类型说明](ucs_api_0024.xml)。
	Provider string `json:"provider"`

	// 所在国家代码，具体代码请参见[国家码](ucs_api_0022.xml)。
	Country string `json:"country"`

	// 所在城市代码，具体代码请参见[城市码](ucs_api_0023.xml)。仅支持中国城市，其他国家无需填写。
	City string `json:"city"`

	// 地域信息。仅在CCE导入集群注册时使用。可通过[获取未注册到UCS的CCE集群](ListManagedClusters.xml)接口的region字段获取。
	Region *string `json:"region,omitempty"`

	// 项目ID信息。仅在CCE导入集群注册时使用。可通过[获取未注册到UCS的CCE集群](ListManagedClusters.xml)接口的projectID字段获取。
	ProjectID *string `json:"projectID,omitempty"`

	// 集群管理类型信息。 取值如下： - grouped：在舰队中纳管的集群 - discrete：未加入舰队的集群
	ManageType string `json:"manageType"`

	Network *NetworkConfig `json:"network,omitempty"`
}

func (o RegisterClusterSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterClusterSpec struct{}"
	}

	return strings.Join([]string{"RegisterClusterSpec", string(data)}, " ")
}
