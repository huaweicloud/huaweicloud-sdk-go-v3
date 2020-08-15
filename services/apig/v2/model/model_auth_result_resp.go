/*
 * APIG
 *
 * API网关（API Gateway）是为开发者、合作伙伴提供的高性能、高可用、高安全的API托管服务，帮助用户轻松构建、管理和发布任意规模的API。
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// 授权结果
type AuthResultResp struct {
	// API名称
	ApiName string `json:"api_name,omitempty"`
	// APP名称
	AppName string `json:"app_name,omitempty"`
	// 授权结果 - SUCCESS：授权成功 - SKIPPED：跳过 - FILAED：授权失败
	Status AuthResultRespStatus `json:"status,omitempty"`
	// 授权失败错误信息
	ErrorMsg string `json:"error_msg,omitempty"`
	// 授权失败错误码
	ErrorCode string `json:"error_code,omitempty"`
}

func (o AuthResultResp) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AuthResultResp", string(data)}, " ")
}

type AuthResultRespStatus struct {
	value string
}

type AuthResultRespStatusEnum struct {
	SUCCESS AuthResultRespStatus
	SKIPPED AuthResultRespStatus
	FAILED  AuthResultRespStatus
}

func GetAuthResultRespStatusEnum() AuthResultRespStatusEnum {
	return AuthResultRespStatusEnum{
		SUCCESS: AuthResultRespStatus{
			value: "SUCCESS",
		},
		SKIPPED: AuthResultRespStatus{
			value: "SKIPPED",
		},
		FAILED: AuthResultRespStatus{
			value: "FAILED",
		},
	}
}

func (c AuthResultRespStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *AuthResultRespStatus) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
