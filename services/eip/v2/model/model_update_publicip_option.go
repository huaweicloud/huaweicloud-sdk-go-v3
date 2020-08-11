/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// 弹性公网IP对象
type UpdatePublicipOption struct {
	// 功能说明：端口id  约束：必须是存在的端口id，如果不带该参数或者值为空时为解除绑定弹性公网IP，如果该端口不存在或端口已绑定弹性公网IP则会提示出错。  和ip_version字段互斥，不能同时更新。
	PortId string `json:"port_id,omitempty"`
	// 功能说明：IP版本信息  取值范围：4和6  4：IPv4  6：IPv6  约束：必须是系统支持的IP版本类型，和port_id互斥，不能同时更新。
	IpVersion int32 `json:"ip_version,omitempty"`
}
