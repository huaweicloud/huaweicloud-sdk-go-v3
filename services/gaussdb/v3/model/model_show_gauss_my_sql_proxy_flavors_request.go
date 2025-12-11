package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowGaussMySqlProxyFlavorsRequest Request Object
type ShowGaussMySqlProxyFlavorsRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 查询的场景
	QueryType *ShowGaussMySqlProxyFlavorsRequestQueryType `json:"query_type,omitempty"`

	// 数据代理的ID, 规格变更场景需要传该参数，过滤掉无法变更的目标规格
	ProxyId *string `json:"proxy_id,omitempty"`
}

func (o ShowGaussMySqlProxyFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGaussMySqlProxyFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ShowGaussMySqlProxyFlavorsRequest", string(data)}, " ")
}

type ShowGaussMySqlProxyFlavorsRequestQueryType struct {
	value string
}

type ShowGaussMySqlProxyFlavorsRequestQueryTypeEnum struct {
	CREATE ShowGaussMySqlProxyFlavorsRequestQueryType
	SCALE  ShowGaussMySqlProxyFlavorsRequestQueryType
}

func GetShowGaussMySqlProxyFlavorsRequestQueryTypeEnum() ShowGaussMySqlProxyFlavorsRequestQueryTypeEnum {
	return ShowGaussMySqlProxyFlavorsRequestQueryTypeEnum{
		CREATE: ShowGaussMySqlProxyFlavorsRequestQueryType{
			value: "create",
		},
		SCALE: ShowGaussMySqlProxyFlavorsRequestQueryType{
			value: "scale",
		},
	}
}

func (c ShowGaussMySqlProxyFlavorsRequestQueryType) Value() string {
	return c.value
}

func (c ShowGaussMySqlProxyFlavorsRequestQueryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowGaussMySqlProxyFlavorsRequestQueryType) UnmarshalJSON(b []byte) error {
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
