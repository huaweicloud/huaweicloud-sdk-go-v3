/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Response Object
type ListSecurityGroupRulesResponse struct {
	// 安全组规则对象列表
	SecurityGroupRules []SecurityGroupRule `json:"security_group_rules,omitempty"`
}
