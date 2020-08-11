/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// 
type PostPaidServerEipExtendParam struct {
	// 公网IP的计费模式。  取值范围：  - prePaid-预付费，即包年包月； - postPaid-后付费，即按需付费；  > 说明： >  > 如果bandwidth对象中share_type是WHOLE且id有值，弹性IP只能创建为按需付费的，故该参数传参“prePaid”无效。
	ChargingMode string `json:"chargingMode,omitempty"`
}
