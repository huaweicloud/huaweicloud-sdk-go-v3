package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowInstanceResponse struct {

	// DDM实例ID。
	Id *string `json:"id,omitempty"`

	// DDM实例状态。
	Status *string `json:"status,omitempty"`

	// DDM实例名称。
	Name *string `json:"name,omitempty"`

	// DDM实例创建时间。
	Created *string `json:"created,omitempty"`

	// DDM实例最后更新时间。
	Updated *string `json:"updated,omitempty"`

	// DDM实例可用区名称。
	AvailableZone *string `json:"available_zone,omitempty"`

	// 虚拟私有云的ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// 子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 安全组ID。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// 节点数量。
	NodeCount *int32 `json:"node_count,omitempty"`

	// DDM实例访问地址。
	AccessIp *string `json:"access_ip,omitempty"`

	// DDM实例访问端口。
	AccessPort *string `json:"access_port,omitempty"`

	// 节点状态。
	NodeStatus *string `json:"node_status,omitempty"`

	// cpu个数。
	CoreCount *string `json:"core_count,omitempty"`

	// 内存大小，单位为G。
	RamCapacity *string `json:"ram_capacity,omitempty"`

	// 响应信息，若无异常信息则不返回该参数。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 订单ID。包周期实例的订单ID，按需实例为空。
	OrderId *string `json:"order_id,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 引擎版本号（Core实例版本号）。
	EngineVersion *string `json:"engine_version,omitempty"`

	// 节点信息。
	Nodes          *[]GetDetailfNodesInfo `json:"nodes,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceResponse", string(data)}, " ")
}
