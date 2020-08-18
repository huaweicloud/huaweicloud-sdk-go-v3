/*
 * LiveAPI
 *
 * 直播服务源站所有接口
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Request Object
type ListRecordConfigsRequest struct {
	Domain     string                             `json:"domain"`
	AppName    *string                            `json:"app_name,omitempty"`
	StreamName *string                            `json:"stream_name,omitempty"`
	Page       *int32                             `json:"page,omitempty"`
	Size       *int32                             `json:"size,omitempty"`
	RecordType ListRecordConfigsRequestRecordType `json:"record_type,omitempty"`
}

func (o ListRecordConfigsRequest) String() string {
	data, _ := json.Marshal(o)
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
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
