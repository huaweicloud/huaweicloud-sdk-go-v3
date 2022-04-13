package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DDM实例的信息。
type ShowInstanceBeanResponse struct {
	// DDM实例ID。

	Id string `json:"id"`
	// DDM实例状态。

	Status string `json:"status"`
	// 创建的实例名称。

	Name string `json:"name"`
	// 创建时间，格式为yyyy-mm-dd Thh:mm:ssZ。  其中，T指定某个时间的开始；Z指示 UTC 时间。

	Created string `json:"created"`
	// 更新时间，格式与“created”完全相同。

	Updated string `json:"updated"`
	// 可用区名称

	AvailableZone string `json:"available_zone"`
	// 虚拟私有云的ID。

	VpcId string `json:"vpc_id"`
	// 子网ID。

	SubnetId string `json:"subnet_id"`
	// 安全组ID。

	SecurityGroupId string `json:"security_group_id"`
	// 节点数量。

	NodeCount int32 `json:"node_count"`
	// DDM实例访问地址。

	AccessIp string `json:"access_ip"`
	// DDM实例访问端口。

	AccessPort string `json:"access_port"`
	// cpu个数。

	CoreCount string `json:"core_count"`
	// 内存大小，单位为G。

	RamCapacity string `json:"ram_capacity"`
	// 响应信息，若无异常信息则不返回该参数

	ErrorMsg *string `json:"error_msg,omitempty"`
	// 节点状态。

	NodeStatus string `json:"node_status"`
	// 企业项目ID。

	EnterpriseProjectId string `json:"enterprise_project_id"`
	// 租户在某一region下的project ID。

	ProjectId string `json:"project_id"`
	// 引擎版本号（Core实例版本号）。

	EngineVersion string `json:"engine_version"`
	// 包周期的实例，有订单id。

	OrderId *string `json:"order_id,omitempty"`
}

func (o ShowInstanceBeanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceBeanResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceBeanResponse", string(data)}, " ")
}
