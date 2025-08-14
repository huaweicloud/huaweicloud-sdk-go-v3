package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListRaspEventsRequest Request Object
type ListRaspEventsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// Host Id
	HostId string `json:"host_id"`

	// 查询时间段的起始时间，毫秒级时间戳(ms)
	StartTime int64 `json:"start_time"`

	// 查询时间段的终止时间，毫秒级时间戳(ms)
	EndTime int64 `json:"end_time"`

	// 默认10
	Limit int32 `json:"limit"`

	// 默认0
	Offset int32 `json:"offset"`

	// 应用类型，包含如下1种。   - java ：java类型应用防护。
	AppType *ListRaspEventsRequestAppType `json:"app_type,omitempty"`

	// 告警级别 |- 告警级别，包含如下1种。 - 0 ：Info级别告警 - 1 ：Low级别告警 - 2 ：Medium级别告警 - 3 ：High级别告警 - 4 ：Critical级别告警
	Severity *string `json:"severity,omitempty"`

	// 攻击标识 |- 攻击标识，包含如下6种。 - Attack Success：攻击成功 - Attack Attempt：攻击尝试 - Attack Blocked：攻击被阻断 - Abnormal Behavior：异常行为 - Collapsible Host：主机失陷 - System Vulnerability：系统脆弱性
	AttackTag *string `json:"attack_tag,omitempty"`

	// 防护状态，包含如下2种。   - closed ：未开启。   - opened ：防护中。
	ProtectStatus *ListRaspEventsRequestProtectStatus `json:"protect_status,omitempty"`
}

func (o ListRaspEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRaspEventsRequest struct{}"
	}

	return strings.Join([]string{"ListRaspEventsRequest", string(data)}, " ")
}

type ListRaspEventsRequestAppType struct {
	value string
}

type ListRaspEventsRequestAppTypeEnum struct {
	JAVA ListRaspEventsRequestAppType
}

func GetListRaspEventsRequestAppTypeEnum() ListRaspEventsRequestAppTypeEnum {
	return ListRaspEventsRequestAppTypeEnum{
		JAVA: ListRaspEventsRequestAppType{
			value: "java",
		},
	}
}

func (c ListRaspEventsRequestAppType) Value() string {
	return c.value
}

func (c ListRaspEventsRequestAppType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRaspEventsRequestAppType) UnmarshalJSON(b []byte) error {
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

type ListRaspEventsRequestProtectStatus struct {
	value string
}

type ListRaspEventsRequestProtectStatusEnum struct {
	CLOSED ListRaspEventsRequestProtectStatus
	OPENED ListRaspEventsRequestProtectStatus
}

func GetListRaspEventsRequestProtectStatusEnum() ListRaspEventsRequestProtectStatusEnum {
	return ListRaspEventsRequestProtectStatusEnum{
		CLOSED: ListRaspEventsRequestProtectStatus{
			value: "closed",
		},
		OPENED: ListRaspEventsRequestProtectStatus{
			value: "opened",
		},
	}
}

func (c ListRaspEventsRequestProtectStatus) Value() string {
	return c.value
}

func (c ListRaspEventsRequestProtectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRaspEventsRequestProtectStatus) UnmarshalJSON(b []byte) error {
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
