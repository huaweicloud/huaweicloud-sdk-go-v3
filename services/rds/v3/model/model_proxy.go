package model

import (
	"encoding/json"

	"strings"
)

type Proxy struct {
	// Proxy实例id。
	PoolId string `json:"pool_id"`
	// Proxy实例开启状态。取值范围：open、closed、frozen、 opening 、closing、freezing、unfreezing，分别对应打开、关闭、已冻结、打开中、关闭中、冻结中、解冻中。
	Status string `json:"status"`
	// Proxy读写分离地址。
	Address string `json:"address"`
	// elb模式的虚拟ip信息。
	ElbVip string `json:"elb_vip"`
	// 弹性公网ip信息。
	Eip string `json:"eip"`
	// Proxy端口信息。
	Port string `json:"port"`
	// Proxy实例状态。 取值范围：creating、normal、deleted
	PoolStatus string `json:"pool_status"`
	// 延时阈值，单位是Byte。
	DelayThresholdInKilobytes string `json:"delay_threshold_in_kilobytes"`
	// Proxy实例规格的cpu数
	Cpu string `json:"cpu"`
	// Proxy实例规格的内存数。
	Mem string `json:"mem"`
	// Proxy节点个数
	NodeNum string     `json:"node_num"`
	Nodes   *ProxyNode `json:"nodes"`
	// Proxy主备模式，取值范围：Ha
	Mode string `json:"mode"`
}

func (o Proxy) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Proxy struct{}"
	}

	return strings.Join([]string{"Proxy", string(data)}, " ")
}
