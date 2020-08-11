/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// 弹性云服务器的网络属性。
type ServerAddress struct {
	// IP地址版本。  - “4”：代表IPv4。 - “6”：代表IPv6。
	Version int32 `json:"version"`
	// IP地址。
	Addr string `json:"addr"`
	// IP地址类型。  - fixed：代表私有IP地址。 - floating：代表浮动IP地址。
	OSEXTIPStype string `json:"OS-EXT-IPS:type,omitempty"`
	// MAC地址。
	OSEXTIPSMACmacAddr string `json:"OS-EXT-IPS-MAC:mac_addr,omitempty"`
	// IP地址对应的端口ID。
	OSEXTIPSportId string `json:"OS-EXT-IPS:port_id,omitempty"`
}
