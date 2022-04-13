package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 主机详细信息
type GetHostListInfo struct {
	// 主机ID

	HostId *string `json:"host_id,omitempty"`
	// 主机IP

	HostIp *string `json:"host_ip,omitempty"`
	// 主机名称

	HostName *string `json:"host_name,omitempty"`
	// 主机状态。 uninstall:未安装 running:运行 offline:离线 error:异常 plugin error:插件错误 installing:安装中 install-fail:安装失败 upgrading:升级中 upgrading-transient:升级中 upgrade failed:升级失败 upgrade-fail:升级失败 uninstalling:卸载中 uninstalling-transient:卸载中 authentication error:鉴权失败

	HostStatus *GetHostListInfoHostStatus `json:"host_status,omitempty"`
	// 主机类型。linux:linux类型,windows:windows类型

	HostType *GetHostListInfoHostType `json:"host_type,omitempty"`
	// 主机版本

	HostVersion *string `json:"host_version,omitempty"`
	// 更新时间

	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o GetHostListInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetHostListInfo struct{}"
	}

	return strings.Join([]string{"GetHostListInfo", string(data)}, " ")
}

type GetHostListInfoHostStatus struct {
	value string
}

type GetHostListInfoHostStatusEnum struct {
	UNINSTALL              GetHostListInfoHostStatus
	RUNNING                GetHostListInfoHostStatus
	OFFLINE                GetHostListInfoHostStatus
	ERROR                  GetHostListInfoHostStatus
	PLUGIN_ERROR           GetHostListInfoHostStatus
	INSTALLING             GetHostListInfoHostStatus
	INSTALL_FAIL           GetHostListInfoHostStatus
	UPGRADING              GetHostListInfoHostStatus
	UPGRADING_TRANSIENT    GetHostListInfoHostStatus
	UPGRADE_FAILED         GetHostListInfoHostStatus
	UPGRADE_FAIL           GetHostListInfoHostStatus
	UNINSTALLING           GetHostListInfoHostStatus
	UNINSTALLING_TRANSIENT GetHostListInfoHostStatus
	AUTHENTICATION_ERROR   GetHostListInfoHostStatus
}

func GetGetHostListInfoHostStatusEnum() GetHostListInfoHostStatusEnum {
	return GetHostListInfoHostStatusEnum{
		UNINSTALL: GetHostListInfoHostStatus{
			value: "uninstall",
		},
		RUNNING: GetHostListInfoHostStatus{
			value: "running",
		},
		OFFLINE: GetHostListInfoHostStatus{
			value: "offline",
		},
		ERROR: GetHostListInfoHostStatus{
			value: "error",
		},
		PLUGIN_ERROR: GetHostListInfoHostStatus{
			value: "plugin error",
		},
		INSTALLING: GetHostListInfoHostStatus{
			value: "installing",
		},
		INSTALL_FAIL: GetHostListInfoHostStatus{
			value: "install-fail",
		},
		UPGRADING: GetHostListInfoHostStatus{
			value: "upgrading",
		},
		UPGRADING_TRANSIENT: GetHostListInfoHostStatus{
			value: "upgrading-transient",
		},
		UPGRADE_FAILED: GetHostListInfoHostStatus{
			value: "upgrade failed",
		},
		UPGRADE_FAIL: GetHostListInfoHostStatus{
			value: "upgrade-fail",
		},
		UNINSTALLING: GetHostListInfoHostStatus{
			value: "uninstalling",
		},
		UNINSTALLING_TRANSIENT: GetHostListInfoHostStatus{
			value: "uninstalling-transient",
		},
		AUTHENTICATION_ERROR: GetHostListInfoHostStatus{
			value: "authentication error",
		},
	}
}

func (c GetHostListInfoHostStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetHostListInfoHostStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type GetHostListInfoHostType struct {
	value string
}

type GetHostListInfoHostTypeEnum struct {
	LINUX   GetHostListInfoHostType
	WINDOWS GetHostListInfoHostType
}

func GetGetHostListInfoHostTypeEnum() GetHostListInfoHostTypeEnum {
	return GetHostListInfoHostTypeEnum{
		LINUX: GetHostListInfoHostType{
			value: "linux",
		},
		WINDOWS: GetHostListInfoHostType{
			value: "windows",
		},
	}
}

func (c GetHostListInfoHostType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetHostListInfoHostType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
