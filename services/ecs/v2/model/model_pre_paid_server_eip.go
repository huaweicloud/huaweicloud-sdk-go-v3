/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// 
type PrePaidServerEip struct {
	// 弹性IP地址类型。
	Iptype string `json:"iptype"`
	Bandwidth *PrePaidServerEipBandwidth `json:"bandwidth"`
	Extendparam *PrePaidServerEipExtendParam `json:"extendparam,omitempty"`
}
