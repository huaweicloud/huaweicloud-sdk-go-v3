package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PluginStatusInfo 服务器的插件状态信息
type PluginStatusInfo struct {

	// **参数解释**: 服务器ID **取值范围**: 字符长度1-128
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器类型 **取值范围**: - host：非容器主机 - container：容器主机
	HostType *PluginStatusInfoHostType `json:"host_type,omitempty"`

	// **参数解释**: 服务器私网IP地址 **取值范围**: 字符长度0-128
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 服务器公网IP地址 **取值范围**: 字符长度0-128
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**: 服务器状态 **取值范围**: - ACTIVE：正在运行 - BUILDING：创建中 - ERROR：故障 - SHUTOFF：关机
	HostStatus *PluginStatusInfoHostStatus `json:"host_status,omitempty"`

	// **参数解释**: Agent状态 **取值范围**: - not_installed：未安装 - online：在线 - offline：离线 - install_failed：安装失败 - installing：安装中
	AgentStatus *PluginStatusInfoAgentStatus `json:"agent_status,omitempty"`

	// **参数解释**： agent版本 **取值范围**： 字符长度0-32位
	AgentVersion *string `json:"agent_version,omitempty"`

	// **参数解释**: 服务器的资产重要性 **取值范围**: - important：重要资产 - common：一般资产 - test：测试资产
	AssetValue *PluginStatusInfoAssetValue `json:"asset_value,omitempty"`

	// **参数解释**: 操作系统类型 **取值范围**: - linux：Linux操作系统 - windows：windows操作系统
	OsType *PluginStatusInfoOsType `json:"os_type,omitempty"`

	// **参数解释**: 系统架构 **取值范围**: - x86_64：X86架构 - arm：ARM架构
	OsArch *PluginStatusInfoOsArch `json:"os_arch,omitempty"`

	// **参数解释**: 系统名称 **取值范围**: 不涉及
	OsName *string `json:"os_name,omitempty"`

	// **参数解释**: 操作系统类型 **取值范围**: - linux：Linux操作系统 - windows：windows操作系统 系统版本
	OsVersion *string `json:"os_version,omitempty"`

	// **参数解释**: 插件状态 **取值范围**: - not_installed：未安装 - installing：安装中 - install_fail：安装失败 - starting：启动中 - running：运行中 - start_fail：启动失败 - offline：离线 - stopping：停止中 - stopped：已停止 - updating：更新中 - update_fail：更新失败 - uninstalling：卸载中 - uninstall_fail：卸载失败
	PluginStatus *PluginStatusInfoPluginStatus `json:"plugin_status,omitempty"`

	// **参数解释**: 插件版本 **取值范围**: 不涉及
	PluginVersion *string `json:"plugin_version,omitempty"`

	// **参数解释**: 插件操作失败原因，包括安装失败、启动失败、离线、停止失败、更新失败以及卸载失败 **取值范围**: 不涉及
	StatusDetail *string `json:"status_detail,omitempty"`

	// **参数解释**: 插件安装进度，百分比 **取值范围**: 0-99
	InstallProgress *string `json:"install_progress,omitempty"`

	// **参数解释**: 插件剩余安装时间，单位分钟 **取值范围**: 不涉及
	RemainingTime *string `json:"remaining_time,omitempty"`

	// **参数解释**: 主机防护状态 **取值范围**: - closed：未开启 - opened：防护中
	ProtectStatus *PluginStatusInfoProtectStatus `json:"protect_status,omitempty"`
}

func (o PluginStatusInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginStatusInfo struct{}"
	}

	return strings.Join([]string{"PluginStatusInfo", string(data)}, " ")
}

type PluginStatusInfoHostType struct {
	value string
}

type PluginStatusInfoHostTypeEnum struct {
	HOST      PluginStatusInfoHostType
	CONTAINER PluginStatusInfoHostType
}

func GetPluginStatusInfoHostTypeEnum() PluginStatusInfoHostTypeEnum {
	return PluginStatusInfoHostTypeEnum{
		HOST: PluginStatusInfoHostType{
			value: "host",
		},
		CONTAINER: PluginStatusInfoHostType{
			value: "container",
		},
	}
}

func (c PluginStatusInfoHostType) Value() string {
	return c.value
}

func (c PluginStatusInfoHostType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PluginStatusInfoHostType) UnmarshalJSON(b []byte) error {
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

type PluginStatusInfoHostStatus struct {
	value string
}

type PluginStatusInfoHostStatusEnum struct {
	ACTIVE   PluginStatusInfoHostStatus
	BUILDING PluginStatusInfoHostStatus
	ERROR    PluginStatusInfoHostStatus
	SHUTOFF  PluginStatusInfoHostStatus
}

func GetPluginStatusInfoHostStatusEnum() PluginStatusInfoHostStatusEnum {
	return PluginStatusInfoHostStatusEnum{
		ACTIVE: PluginStatusInfoHostStatus{
			value: "ACTIVE",
		},
		BUILDING: PluginStatusInfoHostStatus{
			value: "BUILDING",
		},
		ERROR: PluginStatusInfoHostStatus{
			value: "ERROR",
		},
		SHUTOFF: PluginStatusInfoHostStatus{
			value: "SHUTOFF",
		},
	}
}

func (c PluginStatusInfoHostStatus) Value() string {
	return c.value
}

func (c PluginStatusInfoHostStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PluginStatusInfoHostStatus) UnmarshalJSON(b []byte) error {
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

type PluginStatusInfoAgentStatus struct {
	value string
}

type PluginStatusInfoAgentStatusEnum struct {
	NOT_INSTALLED  PluginStatusInfoAgentStatus
	ONLINE         PluginStatusInfoAgentStatus
	OFFLINE        PluginStatusInfoAgentStatus
	INSTALL_FAILED PluginStatusInfoAgentStatus
	INSTALLING     PluginStatusInfoAgentStatus
}

func GetPluginStatusInfoAgentStatusEnum() PluginStatusInfoAgentStatusEnum {
	return PluginStatusInfoAgentStatusEnum{
		NOT_INSTALLED: PluginStatusInfoAgentStatus{
			value: "not_installed",
		},
		ONLINE: PluginStatusInfoAgentStatus{
			value: "online",
		},
		OFFLINE: PluginStatusInfoAgentStatus{
			value: "offline",
		},
		INSTALL_FAILED: PluginStatusInfoAgentStatus{
			value: "install_failed",
		},
		INSTALLING: PluginStatusInfoAgentStatus{
			value: "installing",
		},
	}
}

func (c PluginStatusInfoAgentStatus) Value() string {
	return c.value
}

func (c PluginStatusInfoAgentStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PluginStatusInfoAgentStatus) UnmarshalJSON(b []byte) error {
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

type PluginStatusInfoAssetValue struct {
	value string
}

type PluginStatusInfoAssetValueEnum struct {
	COMMON    PluginStatusInfoAssetValue
	IMPORTANT PluginStatusInfoAssetValue
	TEST      PluginStatusInfoAssetValue
}

func GetPluginStatusInfoAssetValueEnum() PluginStatusInfoAssetValueEnum {
	return PluginStatusInfoAssetValueEnum{
		COMMON: PluginStatusInfoAssetValue{
			value: "common",
		},
		IMPORTANT: PluginStatusInfoAssetValue{
			value: "important",
		},
		TEST: PluginStatusInfoAssetValue{
			value: "test",
		},
	}
}

func (c PluginStatusInfoAssetValue) Value() string {
	return c.value
}

func (c PluginStatusInfoAssetValue) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PluginStatusInfoAssetValue) UnmarshalJSON(b []byte) error {
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

type PluginStatusInfoOsType struct {
	value string
}

type PluginStatusInfoOsTypeEnum struct {
	LINUX   PluginStatusInfoOsType
	WINDOWS PluginStatusInfoOsType
}

func GetPluginStatusInfoOsTypeEnum() PluginStatusInfoOsTypeEnum {
	return PluginStatusInfoOsTypeEnum{
		LINUX: PluginStatusInfoOsType{
			value: "Linux",
		},
		WINDOWS: PluginStatusInfoOsType{
			value: "Windows",
		},
	}
}

func (c PluginStatusInfoOsType) Value() string {
	return c.value
}

func (c PluginStatusInfoOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PluginStatusInfoOsType) UnmarshalJSON(b []byte) error {
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

type PluginStatusInfoOsArch struct {
	value string
}

type PluginStatusInfoOsArchEnum struct {
	X86_64 PluginStatusInfoOsArch
	ARM    PluginStatusInfoOsArch
}

func GetPluginStatusInfoOsArchEnum() PluginStatusInfoOsArchEnum {
	return PluginStatusInfoOsArchEnum{
		X86_64: PluginStatusInfoOsArch{
			value: "x86_64",
		},
		ARM: PluginStatusInfoOsArch{
			value: "arm",
		},
	}
}

func (c PluginStatusInfoOsArch) Value() string {
	return c.value
}

func (c PluginStatusInfoOsArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PluginStatusInfoOsArch) UnmarshalJSON(b []byte) error {
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

type PluginStatusInfoPluginStatus struct {
	value string
}

type PluginStatusInfoPluginStatusEnum struct {
	NOT_INSTALLED  PluginStatusInfoPluginStatus
	INSTALLING     PluginStatusInfoPluginStatus
	INSTALL_FAIL   PluginStatusInfoPluginStatus
	STARTING       PluginStatusInfoPluginStatus
	RUNNING        PluginStatusInfoPluginStatus
	START_FAIL     PluginStatusInfoPluginStatus
	OFFLINE        PluginStatusInfoPluginStatus
	STOPPING       PluginStatusInfoPluginStatus
	STOPPED        PluginStatusInfoPluginStatus
	UPDATING       PluginStatusInfoPluginStatus
	UPDATE_FAIL    PluginStatusInfoPluginStatus
	UNINSTALLING   PluginStatusInfoPluginStatus
	UNINSTALL_FAIL PluginStatusInfoPluginStatus
}

func GetPluginStatusInfoPluginStatusEnum() PluginStatusInfoPluginStatusEnum {
	return PluginStatusInfoPluginStatusEnum{
		NOT_INSTALLED: PluginStatusInfoPluginStatus{
			value: "not_installed",
		},
		INSTALLING: PluginStatusInfoPluginStatus{
			value: "installing",
		},
		INSTALL_FAIL: PluginStatusInfoPluginStatus{
			value: "install_fail",
		},
		STARTING: PluginStatusInfoPluginStatus{
			value: "starting",
		},
		RUNNING: PluginStatusInfoPluginStatus{
			value: "running",
		},
		START_FAIL: PluginStatusInfoPluginStatus{
			value: "start_fail",
		},
		OFFLINE: PluginStatusInfoPluginStatus{
			value: "offline",
		},
		STOPPING: PluginStatusInfoPluginStatus{
			value: "stopping",
		},
		STOPPED: PluginStatusInfoPluginStatus{
			value: "stopped",
		},
		UPDATING: PluginStatusInfoPluginStatus{
			value: "updating",
		},
		UPDATE_FAIL: PluginStatusInfoPluginStatus{
			value: "update_fail",
		},
		UNINSTALLING: PluginStatusInfoPluginStatus{
			value: "uninstalling",
		},
		UNINSTALL_FAIL: PluginStatusInfoPluginStatus{
			value: "uninstall_fail",
		},
	}
}

func (c PluginStatusInfoPluginStatus) Value() string {
	return c.value
}

func (c PluginStatusInfoPluginStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PluginStatusInfoPluginStatus) UnmarshalJSON(b []byte) error {
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

type PluginStatusInfoProtectStatus struct {
	value string
}

type PluginStatusInfoProtectStatusEnum struct {
	CLOSED PluginStatusInfoProtectStatus
	OPENED PluginStatusInfoProtectStatus
}

func GetPluginStatusInfoProtectStatusEnum() PluginStatusInfoProtectStatusEnum {
	return PluginStatusInfoProtectStatusEnum{
		CLOSED: PluginStatusInfoProtectStatus{
			value: "closed",
		},
		OPENED: PluginStatusInfoProtectStatus{
			value: "opened",
		},
	}
}

func (c PluginStatusInfoProtectStatus) Value() string {
	return c.value
}

func (c PluginStatusInfoProtectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PluginStatusInfoProtectStatus) UnmarshalJSON(b []byte) error {
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
