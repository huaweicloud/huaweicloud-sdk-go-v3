package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SwitchAuditDbRequest struct {

	// 数据库ID,可在查询数据库列表接口的ID字段获取。
	Id string `json:"id"`

	// 开关状态 - ON:开启 - OFF:关闭
	Status SwitchAuditDbRequestStatus `json:"status"`

	// 是否关闭LTS审计,DWS数据库场景使用。若用户未选择关闭LTS审计,则不做操作。 - 1 :是 - 0 或 其它:保持原状
	LtsAuditSwitch *int32 `json:"lts_audit_switch,omitempty"`
}

func (o SwitchAuditDbRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchAuditDbRequest struct{}"
	}

	return strings.Join([]string{"SwitchAuditDbRequest", string(data)}, " ")
}

type SwitchAuditDbRequestStatus struct {
	value string
}

type SwitchAuditDbRequestStatusEnum struct {
	ON  SwitchAuditDbRequestStatus
	OFF SwitchAuditDbRequestStatus
}

func GetSwitchAuditDbRequestStatusEnum() SwitchAuditDbRequestStatusEnum {
	return SwitchAuditDbRequestStatusEnum{
		ON: SwitchAuditDbRequestStatus{
			value: "ON",
		},
		OFF: SwitchAuditDbRequestStatus{
			value: "OFF",
		},
	}
}

func (c SwitchAuditDbRequestStatus) Value() string {
	return c.value
}

func (c SwitchAuditDbRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchAuditDbRequestStatus) UnmarshalJSON(b []byte) error {
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
