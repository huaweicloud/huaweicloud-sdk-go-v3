package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPluginStatusRequest Request Object
type ListPluginStatusRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 插件编码 **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	PluginCode string `json:"plugin_code"`

	// **参数解释**： 插件版本 **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	PluginVersion *string `json:"plugin_version,omitempty"`

	// **参数解释**： 插件状态 **约束限制**: 不涉及 **取值范围**: - not_installed：未安装 - installing：安装中 - install_fail：安装失败 - starting：启动中 - running：运行中 - start_fail：启动失败 - offline：离线 - stopping：停止中 - stopped：已停止 - updating：更新中 - update_fail：更新失败 - uninstalling：卸载中 - uninstall_fail：卸载失败  **默认取值**: 不涉及
	PluginStatus *string `json:"plugin_status,omitempty"`

	// **参数解释**： 服务器名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 服务器ID列表 **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	HostIds *[]string `json:"host_ids,omitempty"`

	// **参数解释**： 服务器状态 **约束限制**: 不涉及 **取值范围**: - ACTIVE：正在运行 - BUILDING：创建中 - ERROR：故障 - SHUTOFF：关机  **默认取值**: 不涉及
	HostStatus *[]ListPluginStatusRequestHostStatus `json:"host_status,omitempty"`

	// **参数解释**： agent状态 **约束限制**: 不涉及 **取值范围**: -not_installed：未安装 -online：在线 -offline：离线 -install_failed：安装失败 -installing：安装中  **默认取值**: 不涉及
	AgentStatus *ListPluginStatusRequestAgentStatus `json:"agent_status,omitempty"`

	// **参数解释**： 主机操作系统 **约束限制**: 不涉及 **取值范围**: - linux：Linux操作系统 - Windows：windows操作系统  **默认取值**: 不涉及
	OsType *string `json:"os_type,omitempty"`

	// **参数解释**： 系统架构 **约束限制**: 不涉及 **取值范围**: - x86_64：X86架构 - arm：ARM架构  **默认取值**: 不涉及
	OsArch *string `json:"os_arch,omitempty"`

	// **参数解释**： 服务器类型 **约束限制**: 不涉及 **取值范围**: - host：非容器主机 - container：容器主机  **默认取值**: 不涉及
	HostType *string `json:"host_type,omitempty"`
}

func (o ListPluginStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPluginStatusRequest struct{}"
	}

	return strings.Join([]string{"ListPluginStatusRequest", string(data)}, " ")
}

type ListPluginStatusRequestHostStatus struct {
	value string
}

type ListPluginStatusRequestHostStatusEnum struct {
	ACTIVE   ListPluginStatusRequestHostStatus
	BUILDING ListPluginStatusRequestHostStatus
	ERROR    ListPluginStatusRequestHostStatus
	SHUTOFF  ListPluginStatusRequestHostStatus
}

func GetListPluginStatusRequestHostStatusEnum() ListPluginStatusRequestHostStatusEnum {
	return ListPluginStatusRequestHostStatusEnum{
		ACTIVE: ListPluginStatusRequestHostStatus{
			value: "ACTIVE",
		},
		BUILDING: ListPluginStatusRequestHostStatus{
			value: "BUILDING",
		},
		ERROR: ListPluginStatusRequestHostStatus{
			value: "ERROR",
		},
		SHUTOFF: ListPluginStatusRequestHostStatus{
			value: "SHUTOFF",
		},
	}
}

func (c ListPluginStatusRequestHostStatus) Value() string {
	return c.value
}

func (c ListPluginStatusRequestHostStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPluginStatusRequestHostStatus) UnmarshalJSON(b []byte) error {
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

type ListPluginStatusRequestAgentStatus struct {
	value string
}

type ListPluginStatusRequestAgentStatusEnum struct {
	NOT_INSTALLED  ListPluginStatusRequestAgentStatus
	ONLINE         ListPluginStatusRequestAgentStatus
	OFFLINE        ListPluginStatusRequestAgentStatus
	INSTALL_FAILED ListPluginStatusRequestAgentStatus
	INSTALLING     ListPluginStatusRequestAgentStatus
}

func GetListPluginStatusRequestAgentStatusEnum() ListPluginStatusRequestAgentStatusEnum {
	return ListPluginStatusRequestAgentStatusEnum{
		NOT_INSTALLED: ListPluginStatusRequestAgentStatus{
			value: "not_installed",
		},
		ONLINE: ListPluginStatusRequestAgentStatus{
			value: "online",
		},
		OFFLINE: ListPluginStatusRequestAgentStatus{
			value: "offline",
		},
		INSTALL_FAILED: ListPluginStatusRequestAgentStatus{
			value: "install_failed",
		},
		INSTALLING: ListPluginStatusRequestAgentStatus{
			value: "installing",
		},
	}
}

func (c ListPluginStatusRequestAgentStatus) Value() string {
	return c.value
}

func (c ListPluginStatusRequestAgentStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPluginStatusRequestAgentStatus) UnmarshalJSON(b []byte) error {
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
