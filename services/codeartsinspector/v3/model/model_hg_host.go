package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type HgHost struct {

	// 主机授权状态: * -1 未知 * 0 连接成功 * 1 不可达 * 2 登录失败
	AuthStatus *HgHostAuthStatus `json:"auth_status,omitempty"`

	// 主机组id
	GroupId *string `json:"group_id,omitempty"`

	// 主机id
	Id *string `json:"id,omitempty"`

	// 主机ip
	Ip *string `json:"ip,omitempty"`

	// 跳板机id
	JumperServerId *string `json:"jumper_server_id,omitempty"`

	// 最后一次扫描的id
	LastScanId *string `json:"last_scan_id,omitempty"`

	LastScanInfo *ScanInfoDetail `json:"last_scan_info,omitempty"`

	// 主机名
	Name *string `json:"name,omitempty"`

	// 主机操作系统类型
	OsType *string `json:"os_type,omitempty"`

	// smb_credential_id
	SmbCredentialId *string `json:"smb_credential_id,omitempty"`

	// ssh授权id
	SshCredentialId *string `json:"ssh_credential_id,omitempty"`
}

func (o HgHost) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HgHost struct{}"
	}

	return strings.Join([]string{"HgHost", string(data)}, " ")
}

type HgHostAuthStatus struct {
	value int32
}

type HgHostAuthStatusEnum struct {
	E_MINUS_1 HgHostAuthStatus
	E_0       HgHostAuthStatus
	E_1       HgHostAuthStatus
	E_2       HgHostAuthStatus
}

func GetHgHostAuthStatusEnum() HgHostAuthStatusEnum {
	return HgHostAuthStatusEnum{
		E_MINUS_1: HgHostAuthStatus{
			value: -1,
		}, E_0: HgHostAuthStatus{
			value: 0,
		}, E_1: HgHostAuthStatus{
			value: 1,
		}, E_2: HgHostAuthStatus{
			value: 2,
		},
	}
}

func (c HgHostAuthStatus) Value() int32 {
	return c.value
}

func (c HgHostAuthStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HgHostAuthStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
