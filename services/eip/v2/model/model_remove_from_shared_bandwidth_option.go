/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// 带宽对象
type RemoveFromSharedBandwidthOption struct {
	// 弹性公网IP从共享带宽移除后，会为此弹性公网IP创建独占带宽进行计费。  此参数表示弹性公网IP从共享带宽移除后，使用的独占带宽的计费类型。（bandwidth/traffic）
	ChargeMode string `json:"charge_mode"`
	// 功能说明：要从共享带宽中移除的弹性公网IP或者IPv6端口信息  约束：WHOLE类型的带宽支持多个弹性公网IP或者IPv6端口，跟租户的配额相关，默认一个共享带宽的配额为20
	PublicipInfo []RemovePublicipInfo `json:"publicip_info"`
	// 弹性公网IP从共享带宽移除后，会为此弹性公网IP创建独占带宽进行计费。  此参数表示弹性公网IP从共享带宽移除后，使用的独占带宽的带宽大小。（M）取值范围：默认为1~2000Mbit/s. 可能因为局点配置不同而不同。也跟带宽的计费模式（bandwidth/traffic）相关。
	Size int32 `json:"size"`
}
