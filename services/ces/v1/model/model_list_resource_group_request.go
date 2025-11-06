package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourceGroupRequest Request Object
type ListResourceGroupRequest struct {

	// 资源分组的名称；长度为1-128，只能包含0-9/a-z/A-Z/_/-或汉字；如：ResourceGroup-Test01。
	GroupName *string `json:"group_name,omitempty"`

	// 资源分组的ID，长度为1-128，只能包含0-9/a-z/A-Z；如：rg16063743652226ew93e64p。
	GroupId *string `json:"group_id,omitempty"`

	// 资源分组健康状态，值可为health、unhealth、no_alarm_rule；health表示健康，unhealth表示不健康，no_alarm_rule表示未配置告警规则
	Status *ListResourceGroupRequestStatus `json:"status,omitempty"`

	// 分页起始值，类型为integer，默认值为0。
	Start *int32 `json:"start,omitempty"`

	// 单次查询的条数限制，取值范围(0,100]，默认值为100， 用于限制结果数据条数。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListResourceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceGroupRequest struct{}"
	}

	return strings.Join([]string{"ListResourceGroupRequest", string(data)}, " ")
}

type ListResourceGroupRequestStatus struct {
	value string
}

type ListResourceGroupRequestStatusEnum struct {
	HEALTH        ListResourceGroupRequestStatus
	UNHEALTH      ListResourceGroupRequestStatus
	NO_ALARM_RULE ListResourceGroupRequestStatus
}

func GetListResourceGroupRequestStatusEnum() ListResourceGroupRequestStatusEnum {
	return ListResourceGroupRequestStatusEnum{
		HEALTH: ListResourceGroupRequestStatus{
			value: "health",
		},
		UNHEALTH: ListResourceGroupRequestStatus{
			value: "unhealth",
		},
		NO_ALARM_RULE: ListResourceGroupRequestStatus{
			value: "no_alarm_rule",
		},
	}
}

func (c ListResourceGroupRequestStatus) Value() string {
	return c.value
}

func (c ListResourceGroupRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourceGroupRequestStatus) UnmarshalJSON(b []byte) error {
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
