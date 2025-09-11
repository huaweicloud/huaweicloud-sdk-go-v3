package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AuditScopeSwitchRequestNew struct {

	// 审计范围ID列表
	Ids []string `json:"ids"`

	// 状态  - OFF： 关闭  - ON： 启用
	Status AuditScopeSwitchRequestNewStatus `json:"status"`
}

func (o AuditScopeSwitchRequestNew) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditScopeSwitchRequestNew struct{}"
	}

	return strings.Join([]string{"AuditScopeSwitchRequestNew", string(data)}, " ")
}

type AuditScopeSwitchRequestNewStatus struct {
	value string
}

type AuditScopeSwitchRequestNewStatusEnum struct {
	OFF AuditScopeSwitchRequestNewStatus
	ON  AuditScopeSwitchRequestNewStatus
}

func GetAuditScopeSwitchRequestNewStatusEnum() AuditScopeSwitchRequestNewStatusEnum {
	return AuditScopeSwitchRequestNewStatusEnum{
		OFF: AuditScopeSwitchRequestNewStatus{
			value: "OFF",
		},
		ON: AuditScopeSwitchRequestNewStatus{
			value: "ON",
		},
	}
}

func (c AuditScopeSwitchRequestNewStatus) Value() string {
	return c.value
}

func (c AuditScopeSwitchRequestNewStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuditScopeSwitchRequestNewStatus) UnmarshalJSON(b []byte) error {
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
