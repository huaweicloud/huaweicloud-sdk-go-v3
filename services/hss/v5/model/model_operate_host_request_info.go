package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type OperateHostRequestInfo struct {

	// **参数解释**： 操作类型 **约束限制**： 不涉及 **取值范围**： - ignore：忽略 - un_ignore：取消忽略  **默认取值**： 不涉及
	OperateType OperateHostRequestInfoOperateType `json:"operate_type"`

	// **参数解释**： 主机ID列表 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	HostIdList []string `json:"host_id_list"`
}

func (o OperateHostRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateHostRequestInfo struct{}"
	}

	return strings.Join([]string{"OperateHostRequestInfo", string(data)}, " ")
}

type OperateHostRequestInfoOperateType struct {
	value string
}

type OperateHostRequestInfoOperateTypeEnum struct {
	IGNORE    OperateHostRequestInfoOperateType
	UN_IGNORE OperateHostRequestInfoOperateType
}

func GetOperateHostRequestInfoOperateTypeEnum() OperateHostRequestInfoOperateTypeEnum {
	return OperateHostRequestInfoOperateTypeEnum{
		IGNORE: OperateHostRequestInfoOperateType{
			value: "ignore",
		},
		UN_IGNORE: OperateHostRequestInfoOperateType{
			value: "un_ignore",
		},
	}
}

func (c OperateHostRequestInfoOperateType) Value() string {
	return c.value
}

func (c OperateHostRequestInfoOperateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OperateHostRequestInfoOperateType) UnmarshalJSON(b []byte) error {
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
