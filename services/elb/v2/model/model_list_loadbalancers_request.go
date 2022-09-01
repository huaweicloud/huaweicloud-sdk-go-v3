package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListLoadbalancersRequest struct {

	// 分页查询中每页的负载均衡器个数
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 分页查询的起始的资源id，表示上一页最后一条查询记录的负载均衡器的id。不指定时表示查询第一页。
	Marker *string `json:"marker,omitempty" xml:"marker"`

	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。
	PageReverse *bool `json:"page_reverse,omitempty" xml:"page_reverse"`

	// 负载均衡器的ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 负载均衡器的描述信息。
	Description *string `json:"description,omitempty" xml:"description"`

	// 负载均衡器的名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 负载均衡器的操作状态。取值范围：可以为ONLINE、OFFLINE、DEGRADED、DISABLED或NO_MONITOR。
	OperatingStatus *string `json:"operating_status,omitempty" xml:"operating_status"`

	// 负载均衡器的配置状态。取值范围：可以为ACTIVE、PENDING_CREATE 或者ERROR。
	ProvisioningStatus *string `json:"provisioning_status,omitempty" xml:"provisioning_status"`

	// 负载均衡器的内网IP。
	VipAddress *string `json:"vip_address,omitempty" xml:"vip_address"`

	// 负载均衡器内网IP对应的端口ID。
	VipPortId *string `json:"vip_port_id,omitempty" xml:"vip_port_id"`

	// 负载均衡器所在的子网ID。
	VipSubnetId *string `json:"vip_subnet_id,omitempty" xml:"vip_subnet_id"`

	// 负载均衡器所在的虚拟私有云ID。
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id"`

	// 企业项目ID。创建负载均衡器时，给负载均衡器绑定企业项目ID。取值范围：最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。“0”表示默认企业项目。 关于企业项目ID的获取及企业项目特性的详细信息，请参见《企业管理用户指南》。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 负载均衡器的管理状态。负载均衡器停用时不再接收流量。取值范围：true：启用负载均衡器；false：停用负载均衡器。
	AdminStateUp *bool `json:"admin_state_up,omitempty" xml:"admin_state_up"`

	// 负载均衡器的后端服务器的IP地址
	MemberAddress *string `json:"member_address,omitempty" xml:"member_address"`

	// 负载均衡器的后端服务器对应的弹性云服务器ID
	MemberDeviceId *string `json:"member_device_id,omitempty" xml:"member_device_id"`
}

func (o ListLoadbalancersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoadbalancersRequest struct{}"
	}

	return strings.Join([]string{"ListLoadbalancersRequest", string(data)}, " ")
}
