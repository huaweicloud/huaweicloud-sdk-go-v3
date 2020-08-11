/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// 
type PostPaidServerEip struct {
	// 弹性IP地址类型。
	Iptype string `json:"iptype"`
	Bandwidth *PostPaidServerEipBandwidth `json:"bandwidth"`
	Extendparam *PostPaidServerEipExtendParam `json:"extendparam,omitempty"`
}
