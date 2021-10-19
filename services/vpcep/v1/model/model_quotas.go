package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建配额接口请求结构体
type Quotas struct {
	// 资源类型。支持根据资源类型过滤查询指 定类型的配额。 ● endpoint_service：终端节点服务 ● endpoint：终端节点

	Type *QuotasType `json:"type,omitempty"`
	// 已创建的资源个数。 取值范围：0~quota数。

	Used *int32 `json:"used,omitempty"`
	// 资源的最大配额数。 取值范围：各类型资源默认配额数的最大 值。

	Quota *int32 `json:"quota,omitempty"`
}

func (o Quotas) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Quotas struct{}"
	}

	return strings.Join([]string{"Quotas", string(data)}, " ")
}

type QuotasType struct {
	value string
}

type QuotasTypeEnum struct {
	ENDPOINT_SERVICE QuotasType
	ENDPOINT         QuotasType
}

func GetQuotasTypeEnum() QuotasTypeEnum {
	return QuotasTypeEnum{
		ENDPOINT_SERVICE: QuotasType{
			value: "endpoint_service",
		},
		ENDPOINT: QuotasType{
			value: "endpoint",
		},
	}
}

func (c QuotasType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *QuotasType) UnmarshalJSON(b []byte) error {
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
