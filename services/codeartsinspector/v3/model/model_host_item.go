package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type HostItem struct {

	// 主机IP
	Ip string `json:"ip"`

	// 主机名称
	Name string `json:"name"`

	// 操作系统类型: * linux * windows
	OsType HostItemOsType `json:"os_type"`

	// 主机组id
	GroupId *string `json:"group_id,omitempty"`

	// linux主机ssh授权登录信息ID
	SshCredentialId *string `json:"ssh_credential_id,omitempty"`

	// linux跳板机信息ID
	JumperServerId *string `json:"jumper_server_id,omitempty"`

	// windows主机smb授权登录信息ID
	SmbCredentialId *string `json:"smb_credential_id,omitempty"`
}

func (o HostItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostItem struct{}"
	}

	return strings.Join([]string{"HostItem", string(data)}, " ")
}

type HostItemOsType struct {
	value string
}

type HostItemOsTypeEnum struct {
	LINUX   HostItemOsType
	WINDOWS HostItemOsType
}

func GetHostItemOsTypeEnum() HostItemOsTypeEnum {
	return HostItemOsTypeEnum{
		LINUX: HostItemOsType{
			value: "linux",
		},
		WINDOWS: HostItemOsType{
			value: "windows",
		},
	}
}

func (c HostItemOsType) Value() string {
	return c.value
}

func (c HostItemOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HostItemOsType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
