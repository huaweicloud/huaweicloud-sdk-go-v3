/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// 
type Privateip struct {
	// 私有IP的状态  - ACTIVE：活动的  - DOWN：不可用
	Status string `json:"status"`
	// 私有IP ID
	Id string `json:"id"`
	// 分配IP的子网标识
	SubnetId string `json:"subnet_id"`
	// 项目ID
	TenantId string `json:"tenant_id"`
	// 私有IP的使用者，空表示未使用 取值范围：network:dhcp，network:router_interface_distributed，compute:xxx(xxx对应具体的az名称，例如compute:aa-bb-cc表示是被aa-bb-cc上的虚拟机使用) 约束：此处的取值范围只是本服务支持的类型，其他类型未做标注
	DeviceOwner string `json:"device_owner"`
	// 申请到的私有IP
	IpAddress string `json:"ip_address"`
}
