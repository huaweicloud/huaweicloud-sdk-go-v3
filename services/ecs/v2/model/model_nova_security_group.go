/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// 
type NovaSecurityGroup struct {
	// 安全组描述信息，长度0-255
	Description string `json:"description"`
	// 安全组ID，UUID格式
	Id string `json:"id"`
	// 安全组名字，长度0-255
	Name string `json:"name"`
	// 租户ID或项目ID
	TenantId string `json:"tenant_id"`
	// 安全组规则列表
	Rules []NovaSecurityGroupCommonRule `json:"rules"`
}
