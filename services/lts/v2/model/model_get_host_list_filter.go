package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 查询主机信息过滤参数
type GetHostListFilter struct {

	// 主机名称列表。可以根据主机名称列表，进行批量过滤。
	HostNameList *[]string `json:"host_name_list,omitempty"`

	// 主机ID列表。可以根据主机IP列表，进行批量过滤。
	HostIpList *[]string `json:"host_ip_list,omitempty"`

	// 主机状态。可以根据主机状态进行过滤。 uninstall:未安装 running:运行 offline:离线 error:异常 plugin error:插件错误 installing:安装中 install-fail:安装失败 upgrading:升级中 upgrading-transient:升级中 upgrade failed:升级失败 upgrade-fail:升级失败 uninstalling:卸载中 uninstalling-transient:卸载中 authentication error:鉴权失败
	HostStatus *GetHostListFilterHostStatus `json:"host_status,omitempty"`

	// 主机版本。可以根据主机版本进行过滤。
	HostVersion *string `json:"host_version,omitempty"`
}

func (o GetHostListFilter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetHostListFilter struct{}"
	}

	return strings.Join([]string{"GetHostListFilter", string(data)}, " ")
}

type GetHostListFilterHostStatus struct {
	value string
}

type GetHostListFilterHostStatusEnum struct {
	UNINSTALL              GetHostListFilterHostStatus
	RUNNING                GetHostListFilterHostStatus
	OFFLINE                GetHostListFilterHostStatus
	ERROR                  GetHostListFilterHostStatus
	PLUGIN_ERROR           GetHostListFilterHostStatus
	INSTALLING             GetHostListFilterHostStatus
	INSTALL_FAIL           GetHostListFilterHostStatus
	UPGRADING              GetHostListFilterHostStatus
	UPGRADING_TRANSIENT    GetHostListFilterHostStatus
	UPGRADE_FAILED         GetHostListFilterHostStatus
	UPGRADE_FAIL           GetHostListFilterHostStatus
	UNINSTALLING           GetHostListFilterHostStatus
	UNINSTALLING_TRANSIENT GetHostListFilterHostStatus
	AUTHENTICATION_ERROR   GetHostListFilterHostStatus
}

func GetGetHostListFilterHostStatusEnum() GetHostListFilterHostStatusEnum {
	return GetHostListFilterHostStatusEnum{
		UNINSTALL: GetHostListFilterHostStatus{
			value: "uninstall",
		},
		RUNNING: GetHostListFilterHostStatus{
			value: "running",
		},
		OFFLINE: GetHostListFilterHostStatus{
			value: "offline",
		},
		ERROR: GetHostListFilterHostStatus{
			value: "error",
		},
		PLUGIN_ERROR: GetHostListFilterHostStatus{
			value: "plugin error",
		},
		INSTALLING: GetHostListFilterHostStatus{
			value: "installing",
		},
		INSTALL_FAIL: GetHostListFilterHostStatus{
			value: "install-fail",
		},
		UPGRADING: GetHostListFilterHostStatus{
			value: "upgrading",
		},
		UPGRADING_TRANSIENT: GetHostListFilterHostStatus{
			value: "upgrading-transient",
		},
		UPGRADE_FAILED: GetHostListFilterHostStatus{
			value: "upgrade failed",
		},
		UPGRADE_FAIL: GetHostListFilterHostStatus{
			value: "upgrade-fail",
		},
		UNINSTALLING: GetHostListFilterHostStatus{
			value: "uninstalling",
		},
		UNINSTALLING_TRANSIENT: GetHostListFilterHostStatus{
			value: "uninstalling-transient",
		},
		AUTHENTICATION_ERROR: GetHostListFilterHostStatus{
			value: "authentication error",
		},
	}
}

func (c GetHostListFilterHostStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetHostListFilterHostStatus) UnmarshalJSON(b []byte) error {
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
