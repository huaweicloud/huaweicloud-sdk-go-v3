/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// 
type NovaSecurityGroupCommonGroup struct {
	// 对端安全组的名称
	Name string `json:"name"`
	// 对端安全组所属租户的租户ID
	TenantId string `json:"tenant_id"`
}
