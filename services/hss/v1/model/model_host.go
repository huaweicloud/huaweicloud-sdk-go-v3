package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Host struct {
	// 云主机id

	AgentId *string `json:"agent_id,omitempty"`
	// 云主机id

	HostId *string `json:"host_id,omitempty"`
	// 云主机名称

	HostName *string `json:"host_name,omitempty"`
	// 云主机私有IP

	HostIp *string `json:"host_ip,omitempty"`
	// 云主机公网IP

	PublicIp *string `json:"public_ip,omitempty"`
	// 所属企业项目名称

	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`
	// 服务器组名称

	GroupName *string `json:"group_name,omitempty"`
	// 服务到期时间

	ExpireTime *int64 `json:"expire_time,omitempty"`
	// 策略组名称

	PolicyGroupName *string `json:"policy_group_name,omitempty"`
	// 云主机状态：正在运行：ACTIVE; 关机：SHUTOFF; 创建中：BUILDING; 故障：ERROR;

	HostStatus *HostHostStatus `json:"host_status,omitempty"`
	// 客户端状态, 未注册：not_register; 在线：online; 离线：offline; 所有状态：all;

	AgentStatus *HostAgentStatus `json:"agent_status,omitempty"`
	// 云主机开通的版本,hss.version.null：无； hss.version.basic：基础班；hss.version.enterprise：企业版；hss.version.premium：旗舰版；hss.version.wtp：网页防篡改版

	Version *HostVersion `json:"version,omitempty"`
	// 防护状态, opened：开启；opened：关闭

	ProtectStatus *HostProtectStatus `json:"protect_status,omitempty"`
	// 系统镜像

	OsImage *string `json:"os_image,omitempty"`
	// 系统类型

	OsType *string `json:"os_type,omitempty"`
	// 操作系统位数

	OsBit *string `json:"os_bit,omitempty"`
	// 云主机安全检测结果：undetect：未检测；clean：无风险；risk：有风险

	DetectResult *HostDetectResult `json:"detect_result,omitempty"`
	// 资产风险个数

	RiskPortNum *int32 `json:"risk_port_num,omitempty"`
	// 漏洞风险个数

	RiskVulNum *int32 `json:"risk_vul_num,omitempty"`
	// 入侵风险个数

	RiskIntrusionNum *int32 `json:"risk_intrusion_num,omitempty"`
	// 基线风险个数

	RiskBaselineNum *int32 `json:"risk_baseline_num,omitempty"`
	// 计费模式：packet_cycle：包年包月；on_demand：按需付费

	ChargingMode *HostChargingMode `json:"charging_mode,omitempty"`
	// 云服务资源实例ID（UUID）

	ResourceId *string `json:"resource_id,omitempty"`
}

func (o Host) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Host struct{}"
	}

	return strings.Join([]string{"Host", string(data)}, " ")
}

type HostHostStatus struct {
	value string
}

type HostHostStatusEnum struct {
	ACTIVEHOST_STATUS_ACTIVE     HostHostStatus
	SHUTOFFHOST_STATUS_SHUTOFF   HostHostStatus
	BUILDINGHOST_STATUS_BUILDING HostHostStatus
	ERRORHOST_STATUS_ERROR       HostHostStatus
}

func GetHostHostStatusEnum() HostHostStatusEnum {
	return HostHostStatusEnum{
		ACTIVEHOST_STATUS_ACTIVE: HostHostStatus{
			value: "ACTIVE#HOST_STATUS_ACTIVE",
		},
		SHUTOFFHOST_STATUS_SHUTOFF: HostHostStatus{
			value: "SHUTOFF#HOST_STATUS_SHUTOFF",
		},
		BUILDINGHOST_STATUS_BUILDING: HostHostStatus{
			value: "BUILDING#HOST_STATUS_BUILDING",
		},
		ERRORHOST_STATUS_ERROR: HostHostStatus{
			value: "ERROR#HOST_STATUS_ERROR",
		},
	}
}

func (c HostHostStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *HostHostStatus) UnmarshalJSON(b []byte) error {
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

type HostAgentStatus struct {
	value string
}

type HostAgentStatusEnum struct {
	UNINSTALLAGENT_STATUS_UNINSTALL HostAgentStatus
	ONLINEAGENT_STATUS_ONLINE       HostAgentStatus
	OFFLINEAGENT_STATUS_OFFLINE     HostAgentStatus
}

func GetHostAgentStatusEnum() HostAgentStatusEnum {
	return HostAgentStatusEnum{
		UNINSTALLAGENT_STATUS_UNINSTALL: HostAgentStatus{
			value: "uninstall#AGENT_STATUS_UNINSTALL",
		},
		ONLINEAGENT_STATUS_ONLINE: HostAgentStatus{
			value: "online#AGENT_STATUS_ONLINE",
		},
		OFFLINEAGENT_STATUS_OFFLINE: HostAgentStatus{
			value: "offline#AGENT_STATUS_OFFLINE",
		},
	}
}

func (c HostAgentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *HostAgentStatus) UnmarshalJSON(b []byte) error {
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

type HostVersion struct {
	value string
}

type HostVersionEnum struct {
	HSS_VERSION_NULLVERSION_NULL             HostVersion
	HSS_VERSION_BASICVERSION_BASIC           HostVersion
	HSS_VERSION_ENTERPRISEVERSION_ENTERPRISE HostVersion
	HSS_VERSION_PREMIUMVERSION_PREMIUM       HostVersion
	HSS_VERSION_WTPVERSION_WTP               HostVersion
}

func GetHostVersionEnum() HostVersionEnum {
	return HostVersionEnum{
		HSS_VERSION_NULLVERSION_NULL: HostVersion{
			value: "hss.version.null#VERSION_NULL",
		},
		HSS_VERSION_BASICVERSION_BASIC: HostVersion{
			value: "hss.version.basic#VERSION_BASIC",
		},
		HSS_VERSION_ENTERPRISEVERSION_ENTERPRISE: HostVersion{
			value: "hss.version.enterprise#VERSION_ENTERPRISE",
		},
		HSS_VERSION_PREMIUMVERSION_PREMIUM: HostVersion{
			value: "hss.version.premium#VERSION_PREMIUM",
		},
		HSS_VERSION_WTPVERSION_WTP: HostVersion{
			value: "hss.version.wtp#VERSION_WTP",
		},
	}
}

func (c HostVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *HostVersion) UnmarshalJSON(b []byte) error {
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

type HostProtectStatus struct {
	value string
}

type HostProtectStatusEnum struct {
	CLOSEDPROTECT_STATUS_CLOSED HostProtectStatus
	OPENEDPROTECT_STATUS_OPENED HostProtectStatus
}

func GetHostProtectStatusEnum() HostProtectStatusEnum {
	return HostProtectStatusEnum{
		CLOSEDPROTECT_STATUS_CLOSED: HostProtectStatus{
			value: "closed#PROTECT_STATUS_CLOSED",
		},
		OPENEDPROTECT_STATUS_OPENED: HostProtectStatus{
			value: "opened#PROTECT_STATUS_OPENED",
		},
	}
}

func (c HostProtectStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *HostProtectStatus) UnmarshalJSON(b []byte) error {
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

type HostDetectResult struct {
	value string
}

type HostDetectResultEnum struct {
	UNDETECTDETECT_RESULT_UNDETECT HostDetectResult
	CLEANDETECT_RESULT_CLEAN       HostDetectResult
	RISKDETECT_RESULT_RISK         HostDetectResult
}

func GetHostDetectResultEnum() HostDetectResultEnum {
	return HostDetectResultEnum{
		UNDETECTDETECT_RESULT_UNDETECT: HostDetectResult{
			value: "undetect#DETECT_RESULT_UNDETECT",
		},
		CLEANDETECT_RESULT_CLEAN: HostDetectResult{
			value: "clean#DETECT_RESULT_CLEAN",
		},
		RISKDETECT_RESULT_RISK: HostDetectResult{
			value: "risk#DETECT_RESULT_RISK",
		},
	}
}

func (c HostDetectResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *HostDetectResult) UnmarshalJSON(b []byte) error {
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

type HostChargingMode struct {
	value string
}

type HostChargingModeEnum struct {
	PACKET_CYCLECHARGING_MODE_PACKET_CYCLE HostChargingMode
	ON_DEMANDCHARGING_MODE_ON_DEMAND       HostChargingMode
}

func GetHostChargingModeEnum() HostChargingModeEnum {
	return HostChargingModeEnum{
		PACKET_CYCLECHARGING_MODE_PACKET_CYCLE: HostChargingMode{
			value: "packet_cycle#CHARGING_MODE_PACKET_CYCLE",
		},
		ON_DEMANDCHARGING_MODE_ON_DEMAND: HostChargingMode{
			value: "on_demand#CHARGING_MODE_ON_DEMAND",
		},
	}
}

func (c HostChargingMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *HostChargingMode) UnmarshalJSON(b []byte) error {
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
