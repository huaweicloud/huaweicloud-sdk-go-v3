package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAlarmRulesRequest Request Object
type ListAlarmRulesRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 告警命名空间，取值范围：SYS.CBR,SYS.RDS,SYS.GaussDB
	Namespace *ListAlarmRulesRequestNamespace `json:"namespace,omitempty"`

	// RegionID
	RegionId *string `json:"region_id,omitempty"`

	// 告警规则ID
	AlarmRuleId *string `json:"alarm_rule_id,omitempty"`

	// 分页起始值，默认值为0。
	Offset *int64 `json:"offset,omitempty"`

	// 单次查询的条数限制,取值范围(0,100]，默认值为100，用于限制结果数据条数。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAlarmRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmRulesRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmRulesRequest", string(data)}, " ")
}

type ListAlarmRulesRequestNamespace struct {
	value string
}

type ListAlarmRulesRequestNamespaceEnum struct {
	SYS_CBR      ListAlarmRulesRequestNamespace
	SYS_RDS      ListAlarmRulesRequestNamespace
	SYS_GAUSS_DB ListAlarmRulesRequestNamespace
}

func GetListAlarmRulesRequestNamespaceEnum() ListAlarmRulesRequestNamespaceEnum {
	return ListAlarmRulesRequestNamespaceEnum{
		SYS_CBR: ListAlarmRulesRequestNamespace{
			value: "SYS.CBR",
		},
		SYS_RDS: ListAlarmRulesRequestNamespace{
			value: "SYS.RDS",
		},
		SYS_GAUSS_DB: ListAlarmRulesRequestNamespace{
			value: "SYS.GaussDB",
		},
	}
}

func (c ListAlarmRulesRequestNamespace) Value() string {
	return c.value
}

func (c ListAlarmRulesRequestNamespace) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlarmRulesRequestNamespace) UnmarshalJSON(b []byte) error {
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
