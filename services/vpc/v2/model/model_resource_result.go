/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// 
type ResourceResult struct {
	// 功能说明：根据type过滤查询指定类型的配额  取值范围：vpc，subnet，securityGroup，securityGroupRule，publicIp，vpn，vpngw，vpcPeer，firewall，shareBandwidth，shareBandwidthIP，loadbalancer，listener，physicalConnect，virtualInterface，vpcContainRoutetable，routetableContainRoutes
	Type string `json:"type"`
	// 功能说明：已创建的资源个数  取值范围：0~quota数
	Used int32 `json:"used"`
	// 功能说明：资源的最大配额数  取值范围：各类型资源默认配额数~Integer最大值
	Quota int32 `json:"quota"`
	// 允许修改的配额最小值
	Min int32 `json:"min"`
}
