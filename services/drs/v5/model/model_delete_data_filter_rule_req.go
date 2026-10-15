package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteDataFilterRuleReq 批量删除数据过滤规则请求体对象，包含待删除的规则ID列表和数据过滤类型。
type DeleteDataFilterRuleReq struct {

	// 待删除的数据过滤规则ID列表，支持批量删除多个规则。对应查询接口响应中的`batch_id`字段。
	RuleIds []int64 `json:"rule_ids"`

	// 数据过滤条件类型，指定要删除的过滤规则所属的配置层级，取值common表示一般过滤条件，config表示高级配置过滤条件
	Type DeleteDataFilterRuleReqType `json:"type"`
}

func (o DeleteDataFilterRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataFilterRuleReq struct{}"
	}

	return strings.Join([]string{"DeleteDataFilterRuleReq", string(data)}, " ")
}

type DeleteDataFilterRuleReqType struct {
	value string
}

type DeleteDataFilterRuleReqTypeEnum struct {
	COMMON DeleteDataFilterRuleReqType
	CONFIG DeleteDataFilterRuleReqType
}

func GetDeleteDataFilterRuleReqTypeEnum() DeleteDataFilterRuleReqTypeEnum {
	return DeleteDataFilterRuleReqTypeEnum{
		COMMON: DeleteDataFilterRuleReqType{
			value: "common",
		},
		CONFIG: DeleteDataFilterRuleReqType{
			value: "config",
		},
	}
}

func (c DeleteDataFilterRuleReqType) Value() string {
	return c.value
}

func (c DeleteDataFilterRuleReqType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteDataFilterRuleReqType) UnmarshalJSON(b []byte) error {
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
