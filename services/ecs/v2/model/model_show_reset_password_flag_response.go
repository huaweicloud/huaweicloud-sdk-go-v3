/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Response Object
type ShowResetPasswordFlagResponse struct {
	// 是否支持重置密码。  - True：支持一键重置密码。  - False：不支持一键重置密码。
	ResetpwdFlag string `json:"resetpwd_flag,omitempty"`
}
