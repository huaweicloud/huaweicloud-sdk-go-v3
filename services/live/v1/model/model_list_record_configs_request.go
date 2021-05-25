package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListRecordConfigsRequest struct {
	// 直播播放域名

	Domain string `json:"domain"`
	// 流应用名称

	AppName *string `json:"app_name,omitempty"`
	// 流名

	StreamName *string `json:"stream_name,omitempty"`
	// 分页编号。 默认为0。

	Page *int32 `json:"page,omitempty"`
	// 每页记录数。 取值范围：1-100。 默认为10。

	Size *int32 `json:"size,omitempty"`
	// 录制类型 configer_record：按照配置录制

	RecordType *ListRecordConfigsRequestRecordType `json:"record_type,omitempty"`
}

func (o ListRecordConfigsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListRecordConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListRecordConfigsRequest", string(data)}, " ")
}

type ListRecordConfigsRequestRecordType struct {
	value string
}

type ListRecordConfigsRequestRecordTypeEnum struct {
	CONFIGER_RECORD ListRecordConfigsRequestRecordType
}

func GetListRecordConfigsRequestRecordTypeEnum() ListRecordConfigsRequestRecordTypeEnum {
	return ListRecordConfigsRequestRecordTypeEnum{
		CONFIGER_RECORD: ListRecordConfigsRequestRecordType{
			value: "configer_record",
		},
	}
}

func (c ListRecordConfigsRequestRecordType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListRecordConfigsRequestRecordType) UnmarshalJSON(b []byte) error {
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
