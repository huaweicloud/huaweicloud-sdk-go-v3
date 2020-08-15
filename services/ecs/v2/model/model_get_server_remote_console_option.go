/*
 * ecs
 *
 * ECS Open API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

type GetServerRemoteConsoleOption struct {
	// 远程登录协议，请将protocol配置为“vnc”。
	Protocol GetServerRemoteConsoleOptionProtocol `json:"protocol"`
	// 远程登录的类型，请将type配置为“novnc”。
	Type GetServerRemoteConsoleOptionType `json:"type"`
}

func (o GetServerRemoteConsoleOption) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"GetServerRemoteConsoleOption", string(data)}, " ")
}

type GetServerRemoteConsoleOptionProtocol struct {
	value string
}

type GetServerRemoteConsoleOptionProtocolEnum struct {
	VNC GetServerRemoteConsoleOptionProtocol
}

func GetGetServerRemoteConsoleOptionProtocolEnum() GetServerRemoteConsoleOptionProtocolEnum {
	return GetServerRemoteConsoleOptionProtocolEnum{
		VNC: GetServerRemoteConsoleOptionProtocol{
			value: "vnc",
		},
	}
}

func (c GetServerRemoteConsoleOptionProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *GetServerRemoteConsoleOptionProtocol) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type GetServerRemoteConsoleOptionType struct {
	value string
}

type GetServerRemoteConsoleOptionTypeEnum struct {
	NOVNC GetServerRemoteConsoleOptionType
}

func GetGetServerRemoteConsoleOptionTypeEnum() GetServerRemoteConsoleOptionTypeEnum {
	return GetServerRemoteConsoleOptionTypeEnum{
		NOVNC: GetServerRemoteConsoleOptionType{
			value: "novnc",
		},
	}
}

func (c GetServerRemoteConsoleOptionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *GetServerRemoteConsoleOptionType) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
