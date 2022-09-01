package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节点信息。
type MysqlInstanceNodeInfo struct {

	// 实例ID。
	Id string `json:"id" xml:"id"`

	// 节点名称。
	Name string `json:"name" xml:"name"`

	// 节点类型，master或slave。
	Type *string `json:"type,omitempty" xml:"type"`

	// 节点状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 数据库端口号。
	Port *int32 `json:"port,omitempty" xml:"port"`

	// 节点的读内网地址。
	PrivateReadIps *[]string `json:"private_read_ips,omitempty" xml:"private_read_ips"`

	Volume *MysqlInstanceNodeVolumeInfo `json:"volume,omitempty" xml:"volume"`

	// 可用区。
	AzCode *string `json:"az_code,omitempty" xml:"az_code"`

	// 实例所在的区域。
	RegionCode *string `json:"region_code,omitempty" xml:"region_code"`

	// 创建时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。说明：创建时返回值为空，数据库实例创建成功后该值不为空。
	Created *string `json:"created,omitempty" xml:"created"`

	// 更新时间，格式与\"created\"字段对应格式完全相同。说明：创建时返回值为空，数据库实例创建成功后该值不为空。
	Updated *string `json:"updated,omitempty" xml:"updated"`

	// 规格ID。
	FlavorId *string `json:"flavor_id,omitempty" xml:"flavor_id"`

	// 规格码。
	FlavorRef *string `json:"flavor_ref,omitempty" xml:"flavor_ref"`

	// 允许的最大连接数。
	MaxConnections *string `json:"max_connections,omitempty" xml:"max_connections"`

	// CPU核数。
	Vcpus *string `json:"vcpus,omitempty" xml:"vcpus"`

	// 内存大小，单位为GB。
	Ram *string `json:"ram,omitempty" xml:"ram"`

	// 是否需要重启使修改的参数生效。
	NeedRestart *bool `json:"need_restart,omitempty" xml:"need_restart"`

	// 主备倒换优先级。
	Priority *int32 `json:"priority,omitempty" xml:"priority"`
}

func (o MysqlInstanceNodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlInstanceNodeInfo struct{}"
	}

	return strings.Join([]string{"MysqlInstanceNodeInfo", string(data)}, " ")
}
