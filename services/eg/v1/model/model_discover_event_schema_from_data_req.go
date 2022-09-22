package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 通过事件数据发现事件模型的请求
type DiscoverEventSchemaFromDataReq struct {

	// 事件模型格式类型
	Format *DiscoverEventSchemaFromDataReqFormat `json:"format,omitempty"`

	// 事件数据内容
	Event string `json:"event"`
}

func (o DiscoverEventSchemaFromDataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiscoverEventSchemaFromDataReq struct{}"
	}

	return strings.Join([]string{"DiscoverEventSchemaFromDataReq", string(data)}, " ")
}

type DiscoverEventSchemaFromDataReqFormat struct {
	value string
}

type DiscoverEventSchemaFromDataReqFormatEnum struct {
	JSON_SCHEMA_DRAFT_6 DiscoverEventSchemaFromDataReqFormat
}

func GetDiscoverEventSchemaFromDataReqFormatEnum() DiscoverEventSchemaFromDataReqFormatEnum {
	return DiscoverEventSchemaFromDataReqFormatEnum{
		JSON_SCHEMA_DRAFT_6: DiscoverEventSchemaFromDataReqFormat{
			value: "JSON_SCHEMA_DRAFT_6",
		},
	}
}

func (c DiscoverEventSchemaFromDataReqFormat) Value() string {
	return c.value
}

func (c DiscoverEventSchemaFromDataReqFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DiscoverEventSchemaFromDataReqFormat) UnmarshalJSON(b []byte) error {
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
