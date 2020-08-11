/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

type ServerInterfaceFixedIp struct {
	// 网卡私网IP信息。
	IpAddress string `json:"ip_address,omitempty"`
	// 网卡私网IP对应子网信息。
	SubnetId string `json:"subnet_id,omitempty"`
}
