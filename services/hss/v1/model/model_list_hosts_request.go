package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListHostsRequest struct {
	Version *ListHostsRequestVersion `json:"version,omitempty"`

	AgentStatus *ListHostsRequestAgentStatus `json:"agent_status,omitempty"`

	HostStatus *ListHostsRequestHostStatus `json:"host_status,omitempty"`

	ProtectStatus *ListHostsRequestProtectStatus `json:"protect_status,omitempty"`

	DetectResult *ListHostsRequestDetectResult `json:"detect_result,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	HostIp *string `json:"host_ip,omitempty"`

	PublicIp *string `json:"public_ip,omitempty"`

	OsType *string `json:"os_type,omitempty"`

	ChargingMode *ListHostsRequestChargingMode `json:"charging_mode,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListHostsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListHostsRequest struct{}"
	}

	return strings.Join([]string{"ListHostsRequest", string(data)}, " ")
}

type ListHostsRequestVersion struct {
	value string
}

type ListHostsRequestVersionEnum struct {
	HSS_VERSION_NULLVERSION_NULL             ListHostsRequestVersion
	HSS_VERSION_BASICVERSION_BASIC           ListHostsRequestVersion
	HSS_VERSION_ENTERPRISEVERSION_ENTERPRISE ListHostsRequestVersion
	HSS_VERSION_PREMIUMVERSION_PREMIUM       ListHostsRequestVersion
	HSS_VERSION_WTPVERSION_WTP               ListHostsRequestVersion
}

func GetListHostsRequestVersionEnum() ListHostsRequestVersionEnum {
	return ListHostsRequestVersionEnum{
		HSS_VERSION_NULLVERSION_NULL: ListHostsRequestVersion{
			value: "hss.version.null#VERSION_NULL",
		},
		HSS_VERSION_BASICVERSION_BASIC: ListHostsRequestVersion{
			value: "hss.version.basic#VERSION_BASIC",
		},
		HSS_VERSION_ENTERPRISEVERSION_ENTERPRISE: ListHostsRequestVersion{
			value: "hss.version.enterprise#VERSION_ENTERPRISE",
		},
		HSS_VERSION_PREMIUMVERSION_PREMIUM: ListHostsRequestVersion{
			value: "hss.version.premium#VERSION_PREMIUM",
		},
		HSS_VERSION_WTPVERSION_WTP: ListHostsRequestVersion{
			value: "hss.version.wtp#VERSION_WTP",
		},
	}
}

func (c ListHostsRequestVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListHostsRequestVersion) UnmarshalJSON(b []byte) error {
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

type ListHostsRequestAgentStatus struct {
	value string
}

type ListHostsRequestAgentStatusEnum struct {
	UNINSTALLAGENT_STATUS_UNINSTALL ListHostsRequestAgentStatus
	ONLINEAGENT_STATUS_ONLINE       ListHostsRequestAgentStatus
	OFFLINEAGENT_STATUS_OFFLINE     ListHostsRequestAgentStatus
}

func GetListHostsRequestAgentStatusEnum() ListHostsRequestAgentStatusEnum {
	return ListHostsRequestAgentStatusEnum{
		UNINSTALLAGENT_STATUS_UNINSTALL: ListHostsRequestAgentStatus{
			value: "uninstall#AGENT_STATUS_UNINSTALL",
		},
		ONLINEAGENT_STATUS_ONLINE: ListHostsRequestAgentStatus{
			value: "online#AGENT_STATUS_ONLINE",
		},
		OFFLINEAGENT_STATUS_OFFLINE: ListHostsRequestAgentStatus{
			value: "offline#AGENT_STATUS_OFFLINE",
		},
	}
}

func (c ListHostsRequestAgentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListHostsRequestAgentStatus) UnmarshalJSON(b []byte) error {
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

type ListHostsRequestHostStatus struct {
	value string
}

type ListHostsRequestHostStatusEnum struct {
	ACTIVEHOST_STATUS_ACTIVE     ListHostsRequestHostStatus
	SHUTOFFHOST_STATUS_SHUTOFF   ListHostsRequestHostStatus
	BUILDINGHOST_STATUS_BUILDING ListHostsRequestHostStatus
	ERRORHOST_STATUS_ERROR       ListHostsRequestHostStatus
}

func GetListHostsRequestHostStatusEnum() ListHostsRequestHostStatusEnum {
	return ListHostsRequestHostStatusEnum{
		ACTIVEHOST_STATUS_ACTIVE: ListHostsRequestHostStatus{
			value: "active#HOST_STATUS_ACTIVE",
		},
		SHUTOFFHOST_STATUS_SHUTOFF: ListHostsRequestHostStatus{
			value: "shutoff#HOST_STATUS_SHUTOFF",
		},
		BUILDINGHOST_STATUS_BUILDING: ListHostsRequestHostStatus{
			value: "building#HOST_STATUS_BUILDING",
		},
		ERRORHOST_STATUS_ERROR: ListHostsRequestHostStatus{
			value: "error#HOST_STATUS_ERROR",
		},
	}
}

func (c ListHostsRequestHostStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListHostsRequestHostStatus) UnmarshalJSON(b []byte) error {
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

type ListHostsRequestProtectStatus struct {
	value string
}

type ListHostsRequestProtectStatusEnum struct {
	CLOSEDPROTECT_STATUS_CLOSED ListHostsRequestProtectStatus
	OPENEDPROTECT_STATUS_OPENED ListHostsRequestProtectStatus
}

func GetListHostsRequestProtectStatusEnum() ListHostsRequestProtectStatusEnum {
	return ListHostsRequestProtectStatusEnum{
		CLOSEDPROTECT_STATUS_CLOSED: ListHostsRequestProtectStatus{
			value: "closed#PROTECT_STATUS_CLOSED",
		},
		OPENEDPROTECT_STATUS_OPENED: ListHostsRequestProtectStatus{
			value: "opened#PROTECT_STATUS_OPENED",
		},
	}
}

func (c ListHostsRequestProtectStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListHostsRequestProtectStatus) UnmarshalJSON(b []byte) error {
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

type ListHostsRequestDetectResult struct {
	value string
}

type ListHostsRequestDetectResultEnum struct {
	UNDETECTDETECT_RESULT_UNDETECT ListHostsRequestDetectResult
	CLEANDETECT_RESULT_CLEAN       ListHostsRequestDetectResult
	RISKDETECT_RESULT_RISK         ListHostsRequestDetectResult
}

func GetListHostsRequestDetectResultEnum() ListHostsRequestDetectResultEnum {
	return ListHostsRequestDetectResultEnum{
		UNDETECTDETECT_RESULT_UNDETECT: ListHostsRequestDetectResult{
			value: "undetect#DETECT_RESULT_UNDETECT",
		},
		CLEANDETECT_RESULT_CLEAN: ListHostsRequestDetectResult{
			value: "clean#DETECT_RESULT_CLEAN",
		},
		RISKDETECT_RESULT_RISK: ListHostsRequestDetectResult{
			value: "risk#DETECT_RESULT_RISK",
		},
	}
}

func (c ListHostsRequestDetectResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListHostsRequestDetectResult) UnmarshalJSON(b []byte) error {
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

type ListHostsRequestChargingMode struct {
	value string
}

type ListHostsRequestChargingModeEnum struct {
	PACKET_CYCLECHARGING_MODE_PACKET_CYCLE ListHostsRequestChargingMode
	ON_DEMANDCHARGING_MODE_ON_DEMAND       ListHostsRequestChargingMode
}

func GetListHostsRequestChargingModeEnum() ListHostsRequestChargingModeEnum {
	return ListHostsRequestChargingModeEnum{
		PACKET_CYCLECHARGING_MODE_PACKET_CYCLE: ListHostsRequestChargingMode{
			value: "packet_cycle#CHARGING_MODE_PACKET_CYCLE",
		},
		ON_DEMANDCHARGING_MODE_ON_DEMAND: ListHostsRequestChargingMode{
			value: "on_demand#CHARGING_MODE_ON_DEMAND",
		},
	}
}

func (c ListHostsRequestChargingMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListHostsRequestChargingMode) UnmarshalJSON(b []byte) error {
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
