package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAlarmsRequest Request Object
type ListAlarmsRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 告警命名空间，取值范围：SYS.CBR,SYS.RDS,SYS.GaussDB
	Namespace *ListAlarmsRequestNamespace `json:"namespace,omitempty"`

	// 告警状态，取值范围：ok，alarm，invalid。
	Status *ListAlarmsRequestStatus `json:"status,omitempty"`

	// RegionID
	RegionId *string `json:"region_id,omitempty"`

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 告警规则ID
	AlarmRuleId *string `json:"alarm_rule_id,omitempty"`

	// 查询范围起始时间，例如：2025-03-20T09:31:45Z。
	StartTime *string `json:"start_time,omitempty"`

	// 查询范围结束时间，例如：2025-03-21T09:31:45Z。
	EndTime *string `json:"end_time,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页。
	Marker *string `json:"marker,omitempty"`

	// 单次查询的条数限制,取值范围(0,100]，默认值为10，用于限制结果数据条数。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAlarmsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmsRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmsRequest", string(data)}, " ")
}

type ListAlarmsRequestNamespace struct {
	value string
}

type ListAlarmsRequestNamespaceEnum struct {
	SYS_CBR      ListAlarmsRequestNamespace
	SYS_RDS      ListAlarmsRequestNamespace
	SYS_GAUSS_DB ListAlarmsRequestNamespace
}

func GetListAlarmsRequestNamespaceEnum() ListAlarmsRequestNamespaceEnum {
	return ListAlarmsRequestNamespaceEnum{
		SYS_CBR: ListAlarmsRequestNamespace{
			value: "SYS.CBR",
		},
		SYS_RDS: ListAlarmsRequestNamespace{
			value: "SYS.RDS",
		},
		SYS_GAUSS_DB: ListAlarmsRequestNamespace{
			value: "SYS.GaussDB",
		},
	}
}

func (c ListAlarmsRequestNamespace) Value() string {
	return c.value
}

func (c ListAlarmsRequestNamespace) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlarmsRequestNamespace) UnmarshalJSON(b []byte) error {
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

type ListAlarmsRequestStatus struct {
	value string
}

type ListAlarmsRequestStatusEnum struct {
	OK      ListAlarmsRequestStatus
	ALARM   ListAlarmsRequestStatus
	INVALID ListAlarmsRequestStatus
}

func GetListAlarmsRequestStatusEnum() ListAlarmsRequestStatusEnum {
	return ListAlarmsRequestStatusEnum{
		OK: ListAlarmsRequestStatus{
			value: "ok",
		},
		ALARM: ListAlarmsRequestStatus{
			value: "alarm",
		},
		INVALID: ListAlarmsRequestStatus{
			value: "invalid",
		},
	}
}

func (c ListAlarmsRequestStatus) Value() string {
	return c.value
}

func (c ListAlarmsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlarmsRequestStatus) UnmarshalJSON(b []byte) error {
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
